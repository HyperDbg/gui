// go-cli is the Go reimplementation of hyperdbg-cli.exe. It provides an
// interactive REPL that dispatches command lines through
// api.Debugger.Exec, plus a script mode that runs a file of HyperDbg
// commands line-by-line (.ds) or, in future, a Go script via yaegi
// (.go, Phase C.4.3).
//
// Usage:
//
//	go-cli                                  # interactive REPL
//	go-cli --script find-oep.ds             # run a command script
//	go-cli --script find-oep.go             # run a Go script (yaegi, future)
//	go-cli --connect local --load-vmm --driver Debug\hyperhv.sys
//	go-cli --version
//	go-cli --help
//
// All command logic lives in go-libhyperdbg; the CLI only parses flags,
// constructs an api.Debugger, and routes strings to dbg.Exec. There is
// no global debugger state — the *api.Debugger instance is threaded
// through explicitly so that GUI/MCP layers can reuse the same API.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"github.com/hyperdbg/go-libhyperdbg/symbolparser"
)

// Version constants. The C++ build pulls these from rev.h (CompleteVersion,
// BuildVersion); the Go CLI uses its own constants until the rev module
// is wired up (Phase C.3.6).
const (
	completeVersion = "0.1.0-go"
	buildVersion    = "go-cli"
)

func main() {
	var (
		scriptPath  = flag.String("script", "", "path to a script file (.ds commands or .go for yaegi)")
		connect     = flag.String("connect", "", "connect to target on startup (e.g. 'local' or '<ip> <port>')")
		loadVmm     = flag.Bool("load-vmm", false, "install and start the VMM driver after connecting (requires --driver)")
		driverPath  = flag.String("driver", "", "path to the VMM driver .sys file (used with --load-vmm)")
		symPath     = flag.String("sym-path", "", "DbgHelp symbol search path (enables symbol resolution)")
		showVersion = flag.Bool("version", false, "print version and exit")
		showHelp    = flag.Bool("help", false, "print usage and exit")
	)
	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}
	if *showVersion {
		fmt.Printf("HyperDbg (Go) [version: %s, build: %s]\n", completeVersion, buildVersion)
		return
	}

	// Startup banner — mirrors the printf sequence at the top of the C++
	// hyperdbg-cli main().
	fmt.Printf("HyperDbg Debugger [version: %s, build: %s]\n", completeVersion, buildVersion)
	fmt.Printf("Please visit https://docs.hyperdbg.org for more information...\n")
	fmt.Printf("HyperDbg is released under the GNU Public License v3 (GPLv3).\n\n")

	// Build the debugger options. The symbol resolver must be injected at
	// New() time (the registry captures the output sink), but its Init()
	// runs after the device is open — see find-oep.go for the same pattern.
	opts := []api.Option{api.WithOutput(os.Stdout)}
	var resolver symbolparser.Resolver
	if *symPath != "" {
		resolver = symbolparser.New()
		opts = append(opts, api.WithSymbolResolver(resolver))
	}

	dbg, err := api.New(opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init failed: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = dbg.UnloadVMM(); _ = dbg.UnloadDriver() }()

	// --connect: open the local device (or, in future, a remote target).
	if *connect != "" {
		if err := dbg.Connect(*connect); err != nil {
			fmt.Fprintf(os.Stderr, "connect %q failed: %v\n", *connect, err)
			os.Exit(1)
		}
		fmt.Printf("[*] Connected to %s\n", *connect)
	}

	// --load-vmm: install and start the VMM driver. Requires --driver.
	if *loadVmm {
		if *driverPath == "" {
			fmt.Fprintln(os.Stderr, "err: --load-vmm requires --driver <path>")
			os.Exit(2)
		}
		if err := dbg.LoadDriver(*driverPath); err != nil {
			fmt.Fprintf(os.Stderr, "load-vmm failed: %v\n", err)
			os.Exit(1)
		}
		if err := dbg.InitVMM(); err != nil {
			fmt.Fprintf(os.Stderr, "load-vmm failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[*] VMM loaded (%s)\n", *driverPath)
	}

	// --sym-path: initialise DbgHelp with the given search path. Best
	// effort — if it fails we warn and continue without symbols rather
	// than aborting, since not every workflow needs them.
	if *symPath != "" && resolver != nil {
		if err := resolver.Init(*symPath); err != nil {
			fmt.Fprintf(os.Stderr, "warn: sym-path init failed: %v (continuing without symbols)\n", err)
		} else {
			fmt.Printf("[*] Symbol resolver initialised (sympath=%s)\n", *symPath)
		}
		defer resolver.Close()
	}

	// Dispatch to script mode or the interactive REPL.
	if *scriptPath != "" {
		runner := NewScriptRunner(dbg)
		var runErr error
		switch strings.ToLower(filepath.Ext(*scriptPath)) {
		case ".go":
			runErr = runner.RunGoScript(*scriptPath)
		default:
			// ".ds" or any other extension: treat as a command script.
			runErr = runner.Run(*scriptPath)
		}
		if runErr != nil {
			fmt.Fprintf(os.Stderr, "script error: %v\n", runErr)
			os.Exit(1)
		}
		return
	}

	// Interactive REPL. History persists to ~/.hyperdbg/history.
	repl := NewRepl(
		dbg,
		WithHistoryPath(defaultHistoryPath()),
		WithHistorySize(defaultHistoryMax),
	)
	if err := repl.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "repl error: %v\n", err)
		os.Exit(1)
	}
}

// defaultHistoryPath returns the user's command-history file path
// (~/.hyperdbg/history), or an empty string if the home directory cannot
// be determined. An empty path disables persistence — the REPL still
// keeps history in memory for the session.
func defaultHistoryPath() string {
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	return filepath.Join(home, ".hyperdbg", "history")
}

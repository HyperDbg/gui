package cc

import (
	"testing"

	"github.com/ddkwork/golibrary/stream"
)

type (
	Command struct {
		Cmd   string // todo replace with command kind
		Args  []string
		Usage string
		Demo  []string
	}
)

type Commands struct {
	Debugging []Command
	Extension []Command
	Hwdbg     []Command
	Meta      []Command
}

func TestCommandJson(t *testing.T) {
	commands := &Commands{
		Debugging: []Command{
			{
				Cmd:   "bp",
				Args:  []string{"0x401000"},
				Usage: "set breakpoint at address",
				Demo: []string{
					"bp 0x401000",
					"bp 0x401000,0x401004",
				},
			},
		},
		Extension: []Command{
			{
				Cmd:   "cpuid",
				Args:  []string{""},
				Usage: "get cpuid information",
				Demo:  []string{},
			},
		},
		Hwdbg: []Command{
			{
				Cmd:   "hwclk",
				Args:  []string{},
				Usage: "get hardware clock information",
				Demo:  []string{},
			},
		},
		Meta: []Command{
			{
				Cmd:   "attach",
				Args:  []string{" pid ?"},
				Usage: "attach to a process",
				Demo: []string{
					"pid 1234",
				},
			},
		},
	}
	stream.MarshalJsonToFile(commands, "commands")
}

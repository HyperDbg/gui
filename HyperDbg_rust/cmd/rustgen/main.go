package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func getProjectRoot() string {
	cwd, _ := os.Getwd()
	if filepath.Base(cwd) == "rustgen" {
		return filepath.Dir(filepath.Dir(cwd))
	}
	if filepath.Base(filepath.Dir(cwd)) == "rust-driver" && filepath.Base(cwd) == "kd" {
		return filepath.Dir(cwd)
	}
	return cwd
}

func main() {
	projectRoot := getProjectRoot()
	fmt.Println("=== HyperDbg Code Generator ===")
	fmt.Printf("Project root: %s\n\n", projectRoot)

	fmt.Println("[1/3] Generating Rust types...")
	fmt.Println("  - Output: rust-driver/kd/src/generated/")
	if err := GenerateRustTypes(projectRoot); err != nil {
		fmt.Printf("Error generating Rust types: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n[2/3] Generating WDK bindings...")
	fmt.Println("  - Output: rust-driver/kd/src/generated/")
	if err := GenerateBindgen(projectRoot); err != nil {
		fmt.Printf("Error generating WDK bindings: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n[3/3] Generating MCP server...")
	fmt.Println("  - Server: cmd/mcp/mcp.go")
	if err := GenerateMCP(projectRoot); err != nil {
		fmt.Printf("Error generating MCP server: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n=== All generators completed successfully ===")
}

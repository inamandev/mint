package mint

import (
	"fmt"
	"os"
)

const version = "0.0.1"

const usage = `
mint — a project scaffolder

Usage:
  mint create <project-name>   scaffold a new project
  mint list                 list available templates
  mint version              print mint version

Examples:
  mint create myapp
  mint create myapp --type api --arch ddd
`

func Run() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print(usage)
		os.Exit(0)
	}

	switch args[0] {
	case "create":
		runCreate(args[1:])
	case "list":
		runList()
	case "version", "--version", "-v":
		fmt.Printf("mint version %s\n", version)
	case "help", "--help", "-h":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		fmt.Print(usage)
		os.Exit(1)
	}
}

func runCreate(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: project name required\n\nUsage: mint new <project-name>")
		os.Exit(1)
	}
	projectName := args[0]
	fmt.Printf("scaffolding project: %s\n", projectName)
	// TODO: wire up prompt + scaffold engine
}

func runList() {
	fmt.Println("available templates:")
	fmt.Println("  go-api      REST API (DDD or flat-layered)")
	fmt.Println("  go-grpc     gRPC service")
	fmt.Println("  go-cli      CLI tool")
	fmt.Println("  go-lib      Go library")
	fmt.Println("  go-script   Simple Go script")
}

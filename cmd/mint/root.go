package mint

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/inamandev/mint/internal/prompt"
	tmpl "github.com/inamandev/mint/internal/template"
)

const version = "0.0.2"

const usage = `
mint — a project scaffolder

Usage:
  mint create <project-name>   scaffold a new project
  mint list                    list available templates
  mint version                 print mint version

Examples:
  mint create myapp
`

type Config struct {
	FS fs.FS
}

func Run(c Config) {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Print(usage)
		os.Exit(0)
	}

	switch args[0] {
	case "create":
		runCreate(c.FS, args[1:])
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

func runCreate(templateFS fs.FS, args []string) {
	slug := ""
	if len(args) > 0 {
		slug = args[0]
	} else {
		slug = prompt.Ask("Project slug (used for dir, repo, module):")
	}

	name := prompt.Ask("Project display name:")
	dirName := prompt.Ask("Directory name (default: " + slug + "):")
	if dirName == "" {
		dirName = slug
	}
	module := prompt.Ask("Module path (e.g. github.com/you/" + slug + "):")

	projectType := prompt.Select("Project type:", []string{
		"go-api",
		"go-grpc",
		"go-cli",
		"go-lib",
		"go-script",
	})

	arch := "flat"
	if projectType == "go-api" || projectType == "go-grpc" {
		arch = prompt.Select("Architecture:", []string{"ddd", "flat", "standard"})
	}

	di := "none"
	if prompt.Confirm("Include dependency injection?") {
		di = prompt.Select("DI framework:", []string{"fx", "wire", "dig", "none"})
	}

	features := tmpl.Features{
		Makefile:     prompt.Confirm("Include Makefile?"),
		DevContainer: prompt.Confirm("Include DevContainer?"),
		CI:           prompt.Confirm("Include GitHub Actions CI?"),
		Dockerfile:   prompt.Confirm("Include Dockerfile?"),
		Linter:       prompt.Confirm("Include golangci-lint?"),
		Air:          prompt.Confirm("Include air (hot reload)?"),
	}

	schemaName := projectType + "-" + arch

	fmt.Printf("\n🌿 minting %s...\n", name)

	if err := tmpl.Render(tmpl.RenderOptions{
		FS:         templateFS,
		Language:   "go",
		SchemaName: schemaName,
		Project: tmpl.ProjectData{
			Name:     name,
			Slug:     slug,
			DirName:  dirName,
			Module:   module,
			Arch:     arch,
			DI:       di,
			Features: features,
		},
	}); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ %s created at ./%s\n", name, dirName)
}

func runList() {
	fmt.Println("available templates:")
	fmt.Println("  go-api      REST API (DDD or flat-layered)")
	fmt.Println("  go-grpc     gRPC service")
	fmt.Println("  go-cli      CLI tool")
	fmt.Println("  go-lib      Go library")
	fmt.Println("  go-script   Simple Go script")
}

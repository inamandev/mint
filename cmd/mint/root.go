package mint

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/inamandev/mint/internal/prompt"
	tmpl "github.com/inamandev/mint/internal/template"
	"github.com/inamandev/mint/internal/userconfig"
)

const version = "0.0.3"

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
	cfg, err := userconfig.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: could not load config: %s\n", err)
	}

	// first run — prompt for github username and save
	if cfg.GitHub.Username == "" {
		cfg.GitHub.Username = prompt.Ask("GitHub username:")
		if err := userconfig.Save(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not save config: %s\n", err)
		}
	}

	slug := ""
	if len(args) > 0 {
		slug = args[0]
	} else {
		slug = prompt.Ask("Project slug (used for dir, repo, module):")
	}

	name := prompt.AskDefault("Project display name:", slug)
	dirName := prompt.AskDefault("Directory name:", slug)

	defaultModule := fmt.Sprintf("github.com/%s/%s", cfg.GitHub.Username, slug)
	module := prompt.AskDefault("Module path:", defaultModule)

	projectType := prompt.Select("Project type:", []string{
		"go-api",
		"go-grpc",
		"go-cli",
		"go-lib",
		"go-script",
	})

	defaultArch := cfg.Defaults.Arch
	if defaultArch == "" {
		defaultArch = "ddd"
	}
	arch := "flat"
	if projectType == "go-api" || projectType == "go-grpc" {
		arch = prompt.SelectDefault("Architecture:", []string{"ddd", "flat", "standard"}, defaultArch)
	}

	defaultDI := cfg.Defaults.DI
	if defaultDI == "" {
		defaultDI = "none"
	}
	di := prompt.SelectDefault("DI framework:", []string{"fx", "wire", "dig", "none"}, defaultDI)

	df := cfg.Defaults.Features
	features := tmpl.Features{
		Makefile:     prompt.ConfirmDefault("Include Makefile?", df.Makefile),
		DevContainer: prompt.ConfirmDefault("Include DevContainer?", df.DevContainer),
		CI:           prompt.ConfirmDefault("Include GitHub Actions CI?", df.CI),
		Dockerfile:   prompt.ConfirmDefault("Include Dockerfile?", df.Dockerfile),
		Linter:       prompt.ConfirmDefault("Include golangci-lint?", df.Linter),
		Air:          prompt.ConfirmDefault("Include air (hot reload)?", df.Air),
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

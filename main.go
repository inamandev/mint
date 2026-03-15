package main

import (
	"embed"

	"github.com/inamandev/mint/cmd/mint"
)

// FS holds all embedded template files.
// Templates are bundled into the binary at compile time.
//
//go:embed templates
var TemplateFS embed.FS

func main() {
	mint.Run(mint.Config{FS: TemplateFS})
}

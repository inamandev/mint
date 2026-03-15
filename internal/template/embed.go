package template

import "embed"

// FS holds all embedded template files.
// Templates are bundled into the binary at compile time.
//
//go:embed all:../../templates
var FS embed.FS

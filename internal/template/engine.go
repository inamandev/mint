package template

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

// Schema defines the structure of a template schema JSON file
type Schema struct {
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Required    map[string]string            `json:"required"`
	Optional    map[string]map[string]string `json:"optional"`
}

// Render reads a schema file, renders required files always,
// and optional files based on user feature selections.
func Render(opts RenderOptions) error {
	schemaPath := filepath.Join(opts.Language, "schema", opts.SchemaName+".json")
	schema, err := loadSchema(opts.FS, schemaPath)
	if err != nil {
		return err
	}

	filesDir := filepath.Join(opts.Language, "files")

	// render required files
	for outPath, tmplFile := range schema.Required {
		src := filepath.Join(filesDir, tmplFile)
		dst := filepath.Join(opts.Project.DirName, outPath)
		if err := renderFile(opts.FS, src, dst, opts.Project); err != nil {
			return err
		}
	}

	// render optional files based on features
	for feature, files := range schema.Optional {
		if !featureEnabled(feature, opts.Project.Features) {
			continue
		}
		for outPath, tmplFile := range files {
			src := filepath.Join(filesDir, tmplFile)
			dst := filepath.Join(opts.Project.DirName, outPath)
			if err := renderFile(opts.FS, src, dst, opts.Project); err != nil {
				return err
			}
		}
	}

	return nil
}

// loadSchema reads and parses a schema JSON file from the embedded FS
func loadSchema(fsys fs.FS, schemaPath string) (*Schema, error) {
	content, err := fs.ReadFile(fsys, schemaPath)
	if err != nil {
		return nil, fmt.Errorf("schema not found: %s — %w", schemaPath, err)
	}
	var schema Schema
	if err := json.Unmarshal(content, &schema); err != nil {
		return nil, fmt.Errorf("invalid schema %s — %w", schemaPath, err)
	}
	return &schema, nil
}

// featureEnabled checks if a feature key is enabled in user selections
func featureEnabled(feature string, f Features) bool {
	switch feature {
	case "makefile":
		return f.Makefile
	case "devcontainer":
		return f.DevContainer
	case "ci":
		return f.CI
	case "dockerfile":
		return f.Dockerfile
	case "linter":
		return f.Linter
	case "air":
		return f.Air
	}
	return false
}

// renderFile reads a single template file, runs text/template, writes to disk
func renderFile(fsys fs.FS, srcPath, outPath string, project ProjectData) error {
	content, err := fs.ReadFile(fsys, srcPath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", srcPath, err)
	}
	tmpl, err := template.New(srcPath).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", srcPath, err)
	}
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", outPath, err)
	}
	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outPath, err)
	}
	defer outFile.Close()
	if err := tmpl.Execute(outFile, project); err != nil {
		return fmt.Errorf("failed to render template %s: %w", srcPath, err)
	}
	return nil
}

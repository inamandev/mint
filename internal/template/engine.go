package template

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

// Render walks the embedded template directory for the given project type
// and renders each file to the target output directory.
func Render(templateDir string, outputDir string, data Data) error {
	// get the sub filesystem for the specific template
	subFS, err := fs.Sub(FS, templateDir)
	if err != nil {
		return fmt.Errorf("template not found: %s — %w", templateDir, err)
	}

	return fs.WalkDir(subFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// build the output path
		outPath := filepath.Join(outputDir, path)

		if d.IsDir() {
			return os.MkdirAll(outPath, 0755)
		}

		return renderFile(subFS, path, outPath, data)
	})
}

// renderFile reads a single template file, runs it through text/template
// and writes the result to disk.
func renderFile(subFS fs.FS, srcPath string, outPath string, data Data) error {
	// read template content
	content, err := fs.ReadFile(subFS, srcPath)
	if err != nil {
		return fmt.Errorf("failed to read template file %s: %w", srcPath, err)
	}

	// parse and execute the template
	tmpl, err := template.New(srcPath).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", srcPath, err)
	}

	// ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", outPath, err)
	}

	// create output file
	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outPath, err)
	}
	defer outFile.Close()

	// write rendered content
	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("failed to render template %s: %w", srcPath, err)
	}

	return nil
}

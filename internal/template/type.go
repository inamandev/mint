package template

import "io/fs"

type ProjectData struct {
	Name     string // "My Awesome API"     ← display name
	Slug     string // "my-awesome-api"      ← github repo, go module
	DirName  string // "work/awesome-api"    ← where it gets created on disk
	Module   string // "github.com/you/my-awesome-api"
	Arch     string // "ddd" | "flat" | "standard"
	DI       string // "fx" | "wire" | "dig" | "none"
	Features Features
}

type RenderOptions struct {
	FS         fs.FS
	Language   string
	SchemaName string
	Project    ProjectData
}

// Features represents optional baked-in tooling
type Features struct {
	Makefile     bool
	DevContainer bool
	CI           bool
	Dockerfile   bool
	Linter       bool
	Air          bool
}

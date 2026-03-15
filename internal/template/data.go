package template

// Data holds all variables available inside every template file.
type Data struct {
	ProjectName string // e.g. myapp
	ModulePath  string // e.g. github.com/you/myapp
	ProjectType string // api | grpc | cli | lib | script
	Arch        string // ddd | flat
	DI          string // fx | wire | dig | none
	Features    Features
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

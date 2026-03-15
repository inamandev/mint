# Changelog

## 0.0.3

### Added
- User config system (`~/.mint/config.json`) with `Load()` and `Save()`
- Default-aware prompt functions: `AskDefault`, `SelectDefault`, `ConfirmDefault`
- First-run flow: prompt GitHub username once, save for future runs
- All prompt defaults pre-filled from saved user config
- `go-api-ddd` template schema with required and optional file mappings
- 14 `.tmpl` files for go-api-ddd: main, handler, service, store, domain, logger, config, go.mod, .gitignore, Makefile, Dockerfile, devcontainer, CI, air, linter
- `pkg/logger/logger.go` as required file in schema
- `Config` feature flag in `Features` struct and engine

### Changed
- Moved `internal/config/config.go` from required to optional in go-api-ddd schema
- Fixed `go.mod.tmpl` to use `.Module` instead of `.ModulePath`
- Fixed engine paths to include `templates/` prefix for embedded FS
- Fixed shared `bufio.Reader` in prompt package to prevent input bleeding between calls

## 0.0.2

### Added
- CLI skeleton: `mint create`, `mint list`, `mint version`, `mint help`
- Command router using stdlib `os.Args` + `switch`
- `ProjectData` struct with name/slug/dirname/module separation
- `RenderOptions` struct for engine
- Template engine with schema + flat files approach
- `Render()`, `loadSchema()`, `renderFile()`, `featureEnabled()` in engine
- Prompt functions: `Ask()`, `Select()`, `Confirm()`
- DevContainer setup with go-air and vim local features
- Embedded `fs.FS` interface for templates

## 0.0.1

### Added
- Initial project scaffold
- Go module setup
- Basic Makefile

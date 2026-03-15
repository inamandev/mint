# MINT

A personal project scaffolder. Mint generates opinionated, production-ready project boilerplates so you can skip the setup and start building.

## Features

- **Multiple project types** — REST API, gRPC, CLI, Library, Script
- **Multiple architectures** — DDD (domain-driven) or Flat-layered
- **Optional DI** — choose fx, wire, dig, or none
- **Baked-in tooling** — Makefile, DevContainer, GitHub Actions CI, Dockerfile, air, golangci-lint
- **Interactive wizard** — guided prompts for every option
- **Non-interactive mode** — pass flags for scripting and automation
- **Extensible** — templates for any language, not just Go

---

## Install

```bash
go install github.com/yourusername/mint@latest
```

---

## Usage

### Interactive
```bash
mint new myapp
```

Walks you through a prompt wizard:
```
? Project name:     myapp
? Module path:      github.com/you/myapp
? Project type:     REST API
? Architecture:     DDD
? Include DI?       Yes
? DI framework:     fx
? Features:         Makefile, DevContainer, CI, Dockerfile, Linter
```

### Non-interactive
```bash
mint new myapp --type api --arch ddd --di fx --features makefile,devcontainer,ci
```

### Other commands
```bash
mint list        # list all available templates
mint version     # print mint version
mint help        # show usage
```

---

## Project types

| Type | Description |
|---|---|
| `api` | REST API — DDD or flat-layered |
| `grpc` | gRPC service |
| `cli` | CLI tool |
| `lib` | Go library |
| `script` | Simple Go script |

---

## Generated structure (REST API, DDD)

```
myapp/
├── cmd/server/main.go
├── internal/
│   ├── user/
│   │   ├── domain.go
│   │   ├── service.go
│   │   ├── handler.go
│   │   └── store/
│   └── health/
│       └── handler.go
├── config/
├── pkg/logger/
├── .devcontainer/
├── .github/workflows/ci.yml
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── .golangci.yml
└── go.mod
```

---

## Development

```bash
# clone
git clone https://github.com/yourusername/mint.git
cd mint

# open in devcontainer (recommended)
# VS Code / Zed: reopen in container

# or run locally
make run ARGS="new myapp"
make build
make test
make lint
```

---

## Roadmap

- [ ] REST API template (DDD + flat)
- [ ] gRPC template
- [ ] CLI template
- [ ] Library template
- [ ] Script template
- [ ] Non-interactive flag mode
- [ ] Frontend templates (Next.js, React)

---

## License

MIT

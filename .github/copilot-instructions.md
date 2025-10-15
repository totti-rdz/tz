# TZ - Development Command Tool

A lightweight CLI tool written in Go to streamline development commands across multiple projects with minimal keystrokes.

## Philosophy

- **Minimal keystrokes** - Commands should be as short as possible (prefer single-letter aliases)
- **Universal interface** - Same commands work across all project types
- **Smart defaults** - Auto-detect project type and suggest mappings
- **Fast execution** - Built in Go for instant startup

## Technical Stack

- **Framework**: Cobra (used by kubectl, docker, gh)
- **Language**: Go 1.21+
- **Config storage**: `~/.tz/config.json`
- **Project detection**: Based on current working directory

## Project Structure

```
tz/
├── main.go              # Entry point
├── cmd/                 # Cobra command definitions
│   ├── root.go         # Root command setup
│   ├── install.go      # Install command (alias: i)
│   ├── dev.go          # Dev server (alias: d)
│   ├── test.go         # Test command (alias: t)
│   ├── map.go          # Command mapping
│   └── git.go          # Git shortcuts
├── internal/
│   ├── config/         # Configuration management
│   ├── executor/       # Command execution
│   └── detector/       # Project type detection
└── go.mod
```

## Core Commands

### Development Workflow

- `tz install` / `tz i` - Install dependencies
  - Flags: `-D` for dev dependencies (JavaScript/Node projects)
- `tz dev` / `tz d` - Start development server
- `tz test` / `tz t` - Run tests
- `tz build` / `tz b` - Build project

### Command Mapping

- `tz map <alias> <command>` - Map command for current project
  - Example: `tz map install "npm install"`
  - Example: `tz map i "yarn add"`
  - Example: `tz map dev "npm run start"`

### Git Shortcuts

- `tz fetch` / `tz f` - Git fetch (alias for `git fetch`)
- `tz branch <name>` / `tz br <name>` - Create and checkout branch (alias for `git checkout -b`)
- `tz status` / `tz s` - Git status
- `tz commit <msg>` / `tz c <msg>` - Git commit with message

## Development Guidelines

### Code Style

- Follow standard Go conventions and idioms
- Use Cobra's built-in flag parsing
- Keep commands simple and focused
- Prefer composition over inheritance
- Write unit tests for core logic

### Command Aliases

- Every primary command should have a single-letter alias
- Aliases should be intuitive (install → i, dev → d, test → t)
- Document all aliases in help text

### Configuration

- Store per-project mappings in `~/.tz/config.json`
- Key by absolute project path
- Support global fallback commands
- Auto-create config directory on first run

### Error Handling

- Provide clear, actionable error messages
- Suggest fixes when possible (e.g., "No mapping found. Run: tz map install 'npm install'")
- Exit with appropriate status codes

### Building & Installing

```bash
# Development build
go build -o tz

# Install locally for testing
go install

# Install globally
sudo cp tz /usr/local/bin/
```

## Supported Project Types

The tool should auto-detect and suggest defaults for:

- **JavaScript/TypeScript**: npm, yarn, pnpm, bun
- **Go**: go modules, go test
- **Python**: pip, poetry, pytest
- **Rust**: cargo
- **Java**: maven, gradle
- **Ruby**: bundle

## Future Enhancements

- Shell completion (bash, zsh, fish)
- Interactive project setup wizard
- Command history and suggestions
- Plugin system for custom commands

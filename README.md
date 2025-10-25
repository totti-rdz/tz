# tz - Development Command Tool

A lightweight CLI tool written in Go to streamline development commands across multiple projects with minimal keystrokes. Supports project specific commands and globally defined commands.

## Philosophy

- **Minimal keystrokes** - Single-letter aliases for all commands (`tz i`, `tz d`, `tz t`)
- **Universal interface** - Same commands work across all project types
- **Smart defaults** - Auto-detect project type and suggest commands
- **Zero project pollution** - No config files in your projects, everything stored in `~/.tz/config.json`

## Quick Start

### Installation

```bash
# Clone and build
git clone https://github.com/totti-rdz/tz.git
cd tz
go build -o tz

# Install globally
sudo mv tz /usr/local/bin/
# or
go install
```

### First Time Setup

Navigate to any project and run:

```bash
tz init
```

This will detect your project type and walk you through setting up commands interactively.

## Features

### 🚀 Project-Mapped Commands

Commands that adapt to each project. Just use the same command everywhere:

| Command      | Alias  | Description           | Example Usage                   |
| ------------ | ------ | --------------------- | ------------------------------- |
| `tz install` | `tz i` | Install dependencies  | `tz i express axios`            |
| `tz dev`     | `tz d` | Start dev server      | `tz d --port 8080`              |
| `tz test`    | `tz t` | Run tests             | `tz t user.test.js`             |
| `tz build`   | `tz b` | Build project         | `tz b`                          |
| `tz clear`   | `tz c` | Clear build artifacts | `tz c -a` (includes lock files) |

**All commands accept additional arguments** that get passed to the underlying command!

#### Special Features:

- **`tz i -D`** - Install as dev dependency (works with npm/yarn/pnpm/bun)
- **`tz c -a`** - Clear command + remove lock files (package-lock.json, yarn.lock, etc.)

### 🤖 Smart Auto-Detection

First time running a command in a project? tz detects your project type and suggests the right command:

```bash
$ cd my-node-project
$ tz i

No mapping found for 'install' in this project.
Detected: Node.js project

Run "npm install"? (y/n): y
✓ Saved mapping: install -> "npm install"
# npm install runs...
```

**Supported project types:**

- **Node.js** → npm/yarn/pnpm commands
- **Go** → go mod, go run, go test, go build
- **Python** → pip, pytest, python
- **Rust** → cargo commands
- **Ruby** → bundle commands
- **Java** → mvn commands

### 🎯 Interactive Setup

Set up all commands at once:

```bash
tz init
```

This walks through all commands (install, dev, test, build, clear) and lets you:

- Accept smart suggestions
- Provide custom commands
- Skip commands you don't need

### 🔧 Manual Mapping

Prefer to set commands manually?

```bash
tz map install "yarn install"
tz map dev "npm run start"
tz map test "jest --watch"
tz map build "webpack --mode production"
tz map clear "rm -rf dist build"
```

### ✨ Custom Commands

Create your own project-specific commands with any name:

```bash
tz map docker "docker-compose up"
tz map db "docker-compose --profile infrastructure up"
tz map seed "node scripts/seed.js"
tz map deploy "npm run build && firebase deploy"
```

Then run them directly:

```bash
tz docker           # Runs your docker command
tz db               # Starts database services
tz seed --reset     # Arguments are passed through
```

Custom commands support all the same features as built-in commands (argument passing, etc.).

### 🌎 Global Commands

Create commands that work across **all projects**, not just one:

```bash
tz map --global hello "echo 'hello world'"
tz map -g myglobalscript "echo 'Available everywhere'"
```

Global commands are perfect for personal scripts and tools you use across multiple projects:

```bash
# Set up once
tz map --global lint-all "find . -name '*.js' | xargs eslint"
tz map --global backup "rsync -av . ~/backups/$(basename $PWD)"

# Use anywhere
cd /any/project
tz lint-all        # Works!
tz backup          # Works everywhere!
```

**Priority system:**

- Project-specific commands override global commands
- Built-in commands (`install`, `dev`, `test`, `build`, `clear`) cannot be global

```bash
# Global command
tz map --global docker "docker ps"

# Override in specific project
cd /my/project
tz map docker "docker-compose up"

# In /my/project: runs "docker-compose up"
# Anywhere else: runs "docker ps"
```

### 🎮 Git Shortcuts

Universal git commands that work the same in every project:

|| Command       | Alias     | Description                    | Example                                 |
| ------------- | --------- | ------------------------------ | --------------------------------------- |
| `tz fetch`    | `tz f`    | Git fetch                      | `tz f`                                  |
| `tz branch`   | `tz br`   | Create/checkout branch         | `tz br feature-x`                       |
| `tz branch -` | `tz br -` | Switch to previous branch      | `tz br -`                               |
| `tz status`   | `tz s`    | Git status                     | `tz s`                                  |
| `tz reset`    | `tz r`    | Soft reset commits             | `tz r 2`                                |
| `tz log`      | `tz l`    | View commit history            | `tz l 10`                               |
| `tz clone`    | -         | Clone repo and open in VS Code | `tz clone https://github.com/user/repo` |

#### Advanced Git Features:

**Clone and open projects:**

```bash
tz clone https://github.com/user/repo.git    # Clone and open in VS Code
tz clone git@github.com:user/repo.git        # Works with SSH URLs too
```

**Reset commits safely:**

```bash
tz reset      # Undo last commit (--soft)
tz r 3        # Undo last 3 commits (--soft)
```

**View commit history:**

```bash
tz log        # Show all commits (--oneline)
tz l 5        # Show last 5 commits
tz l 10 -a    # Show last 10 commits (full format)
```

## Examples

### Node.js Project

```bash
cd my-react-app
tz init                    # Set up commands
tz i react-router-dom      # Install package
tz i -D @types/node        # Install dev dependency
tz d                       # Start dev server
tz t                       # Run tests
tz b                       # Build for production
tz c -a                    # Clear build + lock files
```

### Go Project

```bash
cd my-go-api
tz init                    # Auto-detects: go mod download, go run ., etc.
tz i                       # Download dependencies
tz d                       # Run the app
tz t                       # Run tests
tz b                       # Build binary
```

### Mix of Everything

```bash
# Clone a new project
tz clone https://github.com/user/awesome-project.git

# Set up custom commands
tz map docker "docker-compose up"
tz map seed "node scripts/seed.js"

# Use them
tz docker                  # Start containers
tz seed --force            # Seed database

# Regular workflow
tz i axios                 # Install package
tz f                       # Fetch from remote
tz br feature/api          # Create feature branch
tz d                       # Start dev server
tz t api.test.js           # Run specific test
tz s                       # Check git status
tz b                       # Build
tz r                       # Undo last commit
tz r                       # Undo last commit
tz l 5                     # View last 5 commits
```

## Configuration

All mappings are stored in `~/.tz/config.json`:

```json
{
  "global": {
    "hello": "echo 'hello world'",
    "backup": "rsync -av . ~/backups/$(basename $PWD)"
  },
  "projects": {
    "/Users/you/my-project": {
      "install": "npm install",
      "dev": "npm run dev",
      "test": "npm test",
      "build": "npm run build",
      "clear": "rm -rf dist",
      "custom": {
        "docker": "docker-compose up",
        "db": "docker-compose --profile infrastructure up",
        "seed": "node scripts/seed.js"
      }
    },
    "/Users/you/my-go-app": {
      "install": "go mod download",
      "dev": "go run .",
      "test": "go test ./...",
      "build": "go build",
      "clear": "go clean"
    }
  }
}
```

## Why tz?

**Compared to `just` or `make`:**

- ✅ No per-project config files
- ✅ Auto-detection of project types
- ✅ Interactive setup
- ✅ Git shortcuts included
- ✅ Minimal keystrokes (single-letter aliases)

**Compared to npm scripts or project-specific tools:**

- ✅ Works across all languages/frameworks
- ✅ One interface for everything
- ✅ No need to check package.json to remember script names

## Development

Built with:

- **Go 1.21+**
- **Cobra** - CLI framework (used by kubectl, docker, gh)
- **Zero external runtime dependencies**

```bash
# Build
go build -o tz

# Test in a project
./tz init
./tz i express
```

## License

MIT

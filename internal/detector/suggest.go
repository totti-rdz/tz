package detector

// CommandSuggestions holds suggested commands for a project type
type CommandSuggestions struct {
	Install string
	Dev     string
	Test    string
	Build   string
	Clear   string
}

// SuggestCommands returns suggested commands based on project type
func SuggestCommands(projectType ProjectType) *CommandSuggestions {
	switch projectType {
	case NodeJS:
		return &CommandSuggestions{
			Install: "npm install",
			Dev:     "npm run dev",
			Test:    "npm test",
			Build:   "npm run build",
			Clear:   "rm -rf dist",
		}
	case Go:
		return &CommandSuggestions{
			Install: "go mod download",
			Dev:     "go run .",
			Test:    "go test ./...",
			Build:   "go build",
			Clear:   "go clean",
		}
	case Python:
		return &CommandSuggestions{
			Install: "pip install -r requirements.txt",
			Dev:     "python main.py",
			Test:    "pytest",
			Build:   "python -m build",
			Clear:   "rm -rf __pycache__ dist build",
		}
	case Rust:
		return &CommandSuggestions{
			Install: "cargo fetch",
			Dev:     "cargo run",
			Test:    "cargo test",
			Build:   "cargo build",
			Clear:   "cargo clean",
		}
	case Ruby:
		return &CommandSuggestions{
			Install: "bundle install",
			Dev:     "bundle exec rails server",
			Test:    "bundle exec rspec",
			Build:   "bundle exec rake build",
			Clear:   "rm -rf tmp",
		}
	case Java:
		return &CommandSuggestions{
			Install: "mvn install",
			Dev:     "mvn spring-boot:run",
			Test:    "mvn test",
			Build:   "mvn package",
			Clear:   "mvn clean",
		}
	default:
		return nil
	}
}

// GetSuggestion returns a suggested command for a specific command type
func GetSuggestion(projectPath, commandName string) (string, ProjectType) {
	projectType := DetectProjectType(projectPath)
	if projectType == Unknown {
		return "", Unknown
	}

	suggestions := SuggestCommands(projectType)
	if suggestions == nil {
		return "", Unknown
	}

	var suggestion string
	switch commandName {
	case "install":
		suggestion = suggestions.Install
	case "dev":
		suggestion = suggestions.Dev
	case "test":
		suggestion = suggestions.Test
	case "build":
		suggestion = suggestions.Build
	case "clear":
		suggestion = suggestions.Clear
	}

	return suggestion, projectType
}

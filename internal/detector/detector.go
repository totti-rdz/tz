package detector

import (
	"os"
	"path/filepath"
)

// ProjectType represents the detected type of project
type ProjectType string

const (
	NodeJS  ProjectType = "Node.js"
	Go      ProjectType = "Go"
	Python  ProjectType = "Python"
	Rust    ProjectType = "Rust"
	Ruby    ProjectType = "Ruby"
	Java    ProjectType = "Java"
	Unknown ProjectType = "Unknown"
)

// DetectProjectType detects the project type based on marker files in the directory
func DetectProjectType(projectPath string) ProjectType {
	// Check for Node.js
	if fileExists(filepath.Join(projectPath, "package.json")) {
		return NodeJS
	}

	// Check for Go
	if fileExists(filepath.Join(projectPath, "go.mod")) {
		return Go
	}

	// Check for Python
	if fileExists(filepath.Join(projectPath, "pyproject.toml")) ||
		fileExists(filepath.Join(projectPath, "requirements.txt")) ||
		fileExists(filepath.Join(projectPath, "setup.py")) {
		return Python
	}

	// Check for Rust
	if fileExists(filepath.Join(projectPath, "Cargo.toml")) {
		return Rust
	}

	// Check for Ruby
	if fileExists(filepath.Join(projectPath, "Gemfile")) {
		return Ruby
	}

	// Check for Java
	if fileExists(filepath.Join(projectPath, "pom.xml")) ||
		fileExists(filepath.Join(projectPath, "build.gradle")) ||
		fileExists(filepath.Join(projectPath, "build.gradle.kts")) {
		return Java
	}

	return Unknown
}

// fileExists checks if a file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Execute runs a shell command and returns the output or error
func Execute(command string) error {
	if command == "" {
		return fmt.Errorf("empty command")
	}

	// Use shell to execute the command (supports pipes, redirects, etc.)
	cmd := exec.Command("sh", "-c", command)
	
	// Connect to stdout and stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	return nil
}

// ExecuteWithOutput runs a command and returns its output as a string
func ExecuteWithOutput(command string) (string, error) {
	if command == "" {
		return "", fmt.Errorf("empty command")
	}

	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command failed: %w", err)
	}

	return strings.TrimSpace(string(output)), nil
}

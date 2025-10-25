package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/executor"
)

var cloneCmd = &cobra.Command{
	Use:   "clone <repository-url>",
	Short: "Clone a git repository and open it in VS Code",
	Long: `Clone a git repository using git clone and automatically open it in VS Code.

The project name is extracted from the repository URL.

Examples:
  tz clone https://github.com/user/repo.git
  tz clone git@github.com:user/repo.git
  tz clone https://github.com/user/repo`,

	Args: cobra.ExactArgs(1),
	
	RunE: func(cmd *cobra.Command, args []string) error {
		repoURL := args[0]

		// Extract project name from URL
		projectName := extractProjectName(repoURL)
		if projectName == "" {
			return fmt.Errorf("failed to extract project name from URL: %s", repoURL)
		}

		// Clone the repository
		fmt.Printf("Cloning %s...\n", repoURL)
		cloneCmd := fmt.Sprintf("git clone %s", repoURL)
		if err := executor.Execute(cloneCmd); err != nil {
			return fmt.Errorf("failed to clone repository: %w", err)
		}

		// Open in VS Code
		fmt.Printf("Opening %s in VS Code...\n", projectName)
		codeCmd := fmt.Sprintf("code %s", projectName)
		if err := executor.Execute(codeCmd); err != nil {
			// If 'code' command not found, provide helpful message
			if strings.Contains(err.Error(), "command not found") || strings.Contains(err.Error(), "exit status 127") {
				fmt.Printf("\n⚠ VS Code 'code' command not found in PATH.\n")
				fmt.Printf("To enable it:\n")
				fmt.Printf("  1. Open VS Code\n")
				fmt.Printf("  2. Press Cmd+Shift+P\n")
				fmt.Printf("  3. Type 'shell command' and select 'Install code command in PATH'\n")
				fmt.Printf("\nRepository cloned successfully to: %s\n", projectName)
				return nil
			}
			return fmt.Errorf("failed to open in VS Code: %w", err)
		}

		fmt.Printf("✓ Successfully cloned and opened %s\n", projectName)
		return nil
	},
}

// extractProjectName extracts the project name from a git repository URL
func extractProjectName(url string) string {
	// Remove trailing slashes
	url = strings.TrimRight(url, "/")

	// Remove .git suffix if present
	url = strings.TrimSuffix(url, ".git")

	// Handle different URL formats:
	// - https://github.com/user/repo
	// - git@github.com:user/repo
	// - https://github.com/user/repo.git

	var projectName string

	if strings.Contains(url, "://") {
		// HTTP(S) URL format
		parts := strings.Split(url, "/")
		if len(parts) > 0 {
			projectName = parts[len(parts)-1]
		}
	} else if strings.Contains(url, ":") {
		// SSH format (git@github.com:user/repo)
		parts := strings.Split(url, ":")
		if len(parts) == 2 {
			// Get everything after the colon
			pathParts := strings.Split(parts[1], "/")
			if len(pathParts) > 0 {
				projectName = pathParts[len(pathParts)-1]
			}
		}
	} else {
		// Just a path or repo name
		projectName = filepath.Base(url)
	}

	return projectName
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/executor"
)

var rootCmd = &cobra.Command{
	Use:   "tz",
	Short: "A CLI tool to streamline development commands",
	Long: `tz is a lightweight CLI tool that lets you use the same commands 
across all your projects, regardless of language or framework.

Map project-specific commands once, then use simple shortcuts everywhere:
  tz i    - Install dependencies
  tz d    - Start dev server
  tz t    - Run tests
  tz b    - Build project
  tz c    - Clear build artifacts

Universal git shortcuts work everywhere:
  tz f    - Git fetch
  tz s    - Git status
  tz br   - Create and checkout branch`,
	Version: "0.1.0",
}

// Execute runs the root command
func Execute() {
	// Temporarily capture stderr to check for unknown command errors
	err := rootCmd.Execute()
	if err != nil {
		// Check if it's an "unknown command" error
		errStr := err.Error()
		if strings.Contains(errStr, "unknown command") {
			// Extract the command name from the error message
			// Error format: "unknown command \"docker\" for \"tz\""
			parts := strings.Split(errStr, "\"")
			if len(parts) >= 2 {
				commandName := parts[1]
				// Get remaining args from os.Args
				args := os.Args[2:] // Skip program name and command name

				// Try to run as custom command
				if customErr := HandleCustomCommand(commandName, args); customErr != nil {
					fmt.Fprintln(os.Stderr, customErr)
					os.Exit(1)
				}
				return // Success - don't print the original error
			}
		}

		// Other errors - print and exit
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here

	// Silence Cobra's error output so we can handle unknown commands gracefully
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
}

// HandleCustomCommand tries to execute a command as a custom mapped command
func HandleCustomCommand(commandName string, args []string) error {
	// Get current project path
	projectPath, err := config.GetCurrentProjectPath()
	if err != nil {
		return fmt.Errorf("failed to get current project path: %w", err)
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Try to get the custom command
	command, err := cfg.GetCommand(projectPath, commandName)
	if err != nil {
		return fmt.Errorf("unknown command '%s'\n\nTip: Run 'tz map %s \"<your-command>\"' to set it up", commandName, commandName)
	}

	// Append any additional arguments
	if len(args) > 0 {
		command += " " + strings.Join(args, " ")
	}

	// Execute the custom command
	if err := executor.Execute(command); err != nil {
		return err
	}

	return nil
}

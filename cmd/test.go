package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/detector"
	"github.com/totti-rdz/tz/internal/executor"
	"github.com/totti-rdz/tz/internal/prompt"
)

var testCmd = &cobra.Command{
	Use:     "test [args...]",
	Aliases: []string{"t"},
	Short:   "Run tests for current project",
	Long: `Run the test command configured for the current project.

Use 'tz map test <command>' to configure the test command for this project.

Examples:
  tz test              # Run all tests
  tz t                 # Same, using alias
  tz t user.test.js    # Run specific test file
  tz t --watch         # Pass custom arguments`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

		// Get the command mapping
		command, err := cfg.GetCommand(projectPath, "test")
		if err != nil {
			// No mapping found - try auto-detection
			suggestedCmd, projectType := detector.GetSuggestion(projectPath, "test")

			if suggestedCmd == "" || projectType == detector.Unknown {
				return fmt.Errorf("no mapping found for 'test' in this project\n\nTip: Run 'tz map test \"<your-test-command>\"' to set it up")
			}

			// Ask user for confirmation
			if !prompt.ConfirmCommand(string(projectType), "test", suggestedCmd) {
				return fmt.Errorf("cancelled")
			}

			// Save the mapping
			if err := cfg.SetCommand(projectPath, "test", suggestedCmd); err != nil {
				return fmt.Errorf("failed to save mapping: %w", err)
			}

			if err := cfg.Save(); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			command = suggestedCmd
			fmt.Printf("✓ Saved mapping: test -> \"%s\"\n\n", command)
		}

		// Append any additional arguments
		if len(args) > 0 {
			command += " " + strings.Join(args, " ")
		}

		// Execute the command
		if err := executor.Execute(command); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

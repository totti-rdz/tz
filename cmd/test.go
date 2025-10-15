package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/executor"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Run tests for current project",
	Long: `Run the test command configured for the current project.

Use 'tz map test <command>' to configure the test command for this project.

Examples:
  tz test      # Run the configured tests
  tz t         # Same, using alias`,
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
			return fmt.Errorf("%w\n\nTip: Run 'tz map test \"<your-test-command>\"' to set it up", err)
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

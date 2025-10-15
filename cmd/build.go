package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/executor"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Build current project",
	Long: `Run the build command configured for the current project.

Use 'tz map build <command>' to configure the build command for this project.

Examples:
  tz build     # Run the configured build command
  tz b         # Same, using alias`,
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
		command, err := cfg.GetCommand(projectPath, "build")
		if err != nil {
			return fmt.Errorf("%w\n\nTip: Run 'tz map build \"<your-build-command>\"' to set it up", err)
		}

		// Execute the command
		if err := executor.Execute(command); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

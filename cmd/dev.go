package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/detector"
	"github.com/totti-rdz/tz/internal/executor"
	"github.com/totti-rdz/tz/internal/prompt"
)

var devCmd = &cobra.Command{
	Use:     "dev",
	Aliases: []string{"d"},
	Short:   "Run dev server for current project",
	Long: `Run the dev server command configured for the current project.

Use 'tz map dev <command>' to configure the dev command for this project.

Examples:
  tz dev       # Run the configured dev server
  tz d         # Same, using alias`,
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
		command, err := cfg.GetCommand(projectPath, "dev")
		if err != nil {
			// No mapping found - try auto-detection
			suggestedCmd, projectType := detector.GetSuggestion(projectPath, "dev")

			if suggestedCmd == "" || projectType == detector.Unknown {
				return fmt.Errorf("no mapping found for 'dev' in this project\n\nTip: Run 'tz map dev \"<your-dev-command>\"' to set it up")
			}

			// Ask user for confirmation
			if !prompt.ConfirmCommand(string(projectType), "dev", suggestedCmd) {
				return fmt.Errorf("cancelled")
			}

			// Save the mapping
			if err := cfg.SetCommand(projectPath, "dev", suggestedCmd); err != nil {
				return fmt.Errorf("failed to save mapping: %w", err)
			}

			if err := cfg.Save(); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			command = suggestedCmd
			fmt.Printf("✓ Saved mapping: dev -> \"%s\"\n\n", command)
		}

		// Execute the command
		if err := executor.Execute(command); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}

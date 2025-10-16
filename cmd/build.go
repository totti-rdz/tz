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

var buildCmd = &cobra.Command{
	Use:     "build [args...]",
	Aliases: []string{"b"},
	Short:   "Build current project",
	Long: `Run the build command configured for the current project.

Use 'tz map build <command>' to configure the build command for this project.

Examples:
  tz build              # Run the configured build command
  tz b                  # Same, using alias
  tz b --production     # Pass custom arguments`,
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
			// No mapping found - try auto-detection
			suggestedCmd, projectType := detector.GetSuggestion(projectPath, "build")

			if suggestedCmd == "" || projectType == detector.Unknown {
				return fmt.Errorf("no mapping found for 'build' in this project\n\nTip: Run 'tz map build \"<your-build-command>\"' to set it up")
			}

			// Ask user for confirmation
			if !prompt.ConfirmCommand(string(projectType), "build", suggestedCmd) {
				return fmt.Errorf("cancelled")
			}

			// Save the mapping
			if err := cfg.SetCommand(projectPath, "build", suggestedCmd); err != nil {
				return fmt.Errorf("failed to save mapping: %w", err)
			}

			if err := cfg.Save(); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			command = suggestedCmd
			fmt.Printf("✓ Saved mapping: build -> \"%s\"\n\n", command)
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
	rootCmd.AddCommand(buildCmd)
}

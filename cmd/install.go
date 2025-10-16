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

var (
	installDevFlag bool
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Run install command for current project",
	Long: `Run the install command configured for the current project.

Use 'tz map install <command>' to configure the install command for this project.

Examples:
  tz install       # Run the configured install command
  tz i             # Same, using alias
  tz i -D          # Install as dev dependency (npm/yarn/pnpm/bun only)`,
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
		command, err := cfg.GetCommand(projectPath, "install")
		if err != nil {
			// No mapping found - try auto-detection
			suggestedCmd, projectType := detector.GetSuggestion(projectPath, "install")

			if suggestedCmd == "" || projectType == detector.Unknown {
				return fmt.Errorf("no mapping found for 'install' in this project\n\nTip: Run 'tz map install \"<your-install-command>\"' to set it up")
			}

			// Ask user for confirmation
			if !prompt.ConfirmCommand(string(projectType), "install", suggestedCmd) {
				return fmt.Errorf("cancelled")
			}

			// Save the mapping
			if err := cfg.SetCommand(projectPath, "install", suggestedCmd); err != nil {
				return fmt.Errorf("failed to save mapping: %w", err)
			}

			if err := cfg.Save(); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			command = suggestedCmd
			fmt.Printf("✓ Saved mapping: install -> \"%s\"\n\n", command)
		}

		// Handle -D flag for dev dependencies
		if installDevFlag {
			if strings.Contains(command, "npm") || strings.Contains(command, "pnpm") {
				command += " --save-dev"
			} else if strings.Contains(command, "yarn") {
				command += " --dev"
			} else if strings.Contains(command, "bun") {
				command += " -D"
			} else {
				return fmt.Errorf("-D flag is only supported for npm/yarn/pnpm/bun projects\nCurrent command: %s", command)
			}
		}

		// Execute the command
		if err := executor.Execute(command); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&installDevFlag, "dev", "D", false, "Install as dev dependency (npm/yarn/pnpm/bun)")
}

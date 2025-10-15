package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/executor"
)

var (
	clearAllFlag bool
)

var clearCmd = &cobra.Command{
	Use:     "clear",
	Aliases: []string{"c"},
	Short:   "Clear build artifacts for current project",
	Long: `Run the clear command configured for the current project.

Use 'tz map clear <command>' to configure the clear command for this project.

Examples:
  tz clear     # Run the configured clear command
  tz c         # Same, using alias
  tz c -a      # Clear + remove lock files (package-lock.json, yarn.lock, etc.)`,
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
		command, err := cfg.GetCommand(projectPath, "clear")
		if err != nil {
			return fmt.Errorf("%w\n\nTip: Run 'tz map clear \"<your-clear-command>\"' to set it up", err)
		}

		// Execute the mapped clear command
		if err := executor.Execute(command); err != nil {
			return err
		}

		// If -a flag is set, also remove lock files
		if clearAllFlag {
			lockFiles := []string{
				"package-lock.json",
				"yarn.lock",
				"pnpm-lock.yaml",
				"bun.lockb",
				"Gemfile.lock",
				"Cargo.lock",
				"poetry.lock",
			}

			for _, lockFile := range lockFiles {
				lockPath := filepath.Join(projectPath, lockFile)
				if _, err := os.Stat(lockPath); err == nil {
					fmt.Printf("Removing %s...\n", lockFile)
					if err := os.Remove(lockPath); err != nil {
						fmt.Printf("Warning: failed to remove %s: %v\n", lockFile, err)
					}
				}
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
	clearCmd.Flags().BoolVarP(&clearAllFlag, "all", "a", false, "Also remove lock files (package-lock.json, yarn.lock, etc.)")
}

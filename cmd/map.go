package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
)

var (
	globalFlag bool
)

var mapCmd = &cobra.Command{
	Use:   "map <command> <shell-command>",
	Short: "Map a command for the current project or globally",
	Long: `Map a command to a shell command for the current project or globally.

The mapping will be saved in ~/.tz/config.json and associated with 
the current directory (or globally if --global flag is used).

Built-in commands (install, dev, test, build, clear) can only be mapped 
per-project and cannot be set as global commands.

You can also create custom commands with any name you want.

Examples:
  tz map install "npm install"
  tz map dev "npm run start"
  tz map test "go test ./..."
  tz map clear "rm -rf dist"
  tz map docker "docker-compose up"
  tz map seed "node scripts/seed.js"
  tz map --global mouflon "/path/to/mouflon.ts"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		commandName := args[0]
		shellCommand := args[1]

		// Load config
		cfg, err := config.Load()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		// Handle global command mapping
		if globalFlag {
			if err := cfg.SetGlobalCommand(commandName, shellCommand); err != nil {
				return fmt.Errorf("failed to set global command: %w", err)
			}

			// Save config
			if err := cfg.Save(); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			fmt.Printf("✓ Mapped '%s' to '%s' globally\n", commandName, shellCommand)
			return nil
		}

		// Handle project-specific command mapping
		projectPath, err := config.GetCurrentProjectPath()
		if err != nil {
			return fmt.Errorf("failed to get current project path: %w", err)
		}

		// Set the command mapping
		if err := cfg.SetCommand(projectPath, commandName, shellCommand); err != nil {
			return fmt.Errorf("failed to set command: %w", err)
		}

		// Save config
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Printf("✓ Mapped '%s' to '%s' for project:\n  %s\n", commandName, shellCommand, projectPath)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(mapCmd)
	mapCmd.Flags().BoolVarP(&globalFlag, "global", "g", false, "Set command as a global command (not project-specific)")
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
)

var mapCmd = &cobra.Command{
	Use:   "map <command> <shell-command>",
	Short: "Map a command for the current project",
	Long: `Map a command to a shell command for the current project.

The mapping will be saved in ~/.tz/config.json and associated with 
the current directory.

Built-in commands: install, dev, test, build, clear
You can also create custom commands with any name you want.

Examples:
  tz map install "npm install"
  tz map dev "npm run start"
  tz map test "go test ./..."
  tz map clear "rm -rf dist"
  tz map docker "docker-compose up"
  tz map seed "node scripts/seed.js"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		commandName := args[0]
		shellCommand := args[1]

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
}

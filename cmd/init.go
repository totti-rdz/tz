package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/config"
	"github.com/totti-rdz/tz/internal/detector"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize command mappings for current project",
	Long: `Interactive setup to configure all command mappings for the current project.

For each command (install, dev, test, build, clear), tz will:
1. Detect your project type and suggest a command, or
2. Ask you to provide a custom command

You can accept suggestions, provide your own, or skip commands.

Example:
  tz init     # Start interactive setup`,
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

		// Detect project type
		projectType := detector.DetectProjectType(projectPath)

		fmt.Printf("Initializing tz for current project:\n  %s\n\n", projectPath)

		if projectType != detector.Unknown {
			fmt.Printf("Detected: %s project\n\n", projectType)
		} else {
			fmt.Printf("Project type: Unknown (manual configuration required)\n\n")
		}

		reader := bufio.NewReader(os.Stdin)
		commands := []string{"install", "dev", "test", "build", "clear"}

		for _, commandName := range commands {
			// Check if already configured
			existingCmd, err := cfg.GetCommand(projectPath, commandName)
			if err == nil && existingCmd != "" {
				fmt.Printf("✓ '%s' already configured: %s\n", commandName, existingCmd)
				fmt.Print("  Keep this? (y/n): ")
				response, _ := reader.ReadString('\n')
				response = strings.ToLower(strings.TrimSpace(response))
				if response == "y" || response == "yes" || response == "" {
					continue
				}
			}

			// Try to get suggestion
			suggestion, _ := detector.GetSuggestion(projectPath, commandName)

			var selectedCommand string

			if suggestion != "" {
				fmt.Printf("\nCommand: %s\n", commandName)
				fmt.Printf("Suggested: %s\n", suggestion)
				fmt.Print("Accept suggestion? (y/n/custom): ")

				response, _ := reader.ReadString('\n')
				response = strings.ToLower(strings.TrimSpace(response))

				if response == "y" || response == "yes" || response == "" {
					selectedCommand = suggestion
				} else if response == "n" || response == "no" {
					fmt.Printf("Enter command for '%s' (or press Enter to skip): ", commandName)
					customCmd, _ := reader.ReadString('\n')
					customCmd = strings.TrimSpace(customCmd)
					if customCmd != "" {
						selectedCommand = customCmd
					}
				} else {
					// User typed a custom command directly
					selectedCommand = response
				}
			} else {
				// No suggestion available
				fmt.Printf("\nCommand: %s\n", commandName)
				fmt.Printf("Enter command for '%s' (or press Enter to skip): ", commandName)
				customCmd, _ := reader.ReadString('\n')
				customCmd = strings.TrimSpace(customCmd)
				if customCmd != "" {
					selectedCommand = customCmd
				}
			}

			// Save the command if provided
			if selectedCommand != "" {
				if err := cfg.SetCommand(projectPath, commandName, selectedCommand); err != nil {
					return fmt.Errorf("failed to set command: %w", err)
				}
				fmt.Printf("  ✓ Saved: %s -> \"%s\"\n", commandName, selectedCommand)
			} else {
				fmt.Printf("  ⊘ Skipped: %s\n", commandName)
			}
		}

		// Save config
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Printf("\n✓ Configuration complete!\n")
		fmt.Printf("\nYou can now use:\n")
		fmt.Printf("  tz i (install), tz d (dev), tz t (test), tz b (build), tz c (clear)\n")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

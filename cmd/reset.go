package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/executor"
)

var resetCmd = &cobra.Command{
	Use:     "reset [count]",
	Aliases: []string{"r"},
	Short:   "Soft reset commits (git reset --soft HEAD~N)",
	Long: `Soft reset to undo commits while keeping changes staged.

The count parameter specifies how many commits to undo (default: 1).
This always uses --soft flag for safety. Use git directly for hard resets.

Examples:
  tz reset      # Undo last commit (git reset --soft HEAD~1)
  tz r          # Same, using alias
  tz reset 2    # Undo last 2 commits (git reset --soft HEAD~2)
  tz r 3        # Undo last 3 commits`,
	RunE: func(cmd *cobra.Command, args []string) error {
		count := 1 // Default to 1

		// Parse count if provided
		if len(args) > 0 {
			countStr := strings.TrimPrefix(args[0], "-")
			parsed, err := strconv.Atoi(countStr)
			if err != nil {
				return fmt.Errorf("invalid count: %s (must be a number)", args[0])
			}
			if parsed < 1 {
				return fmt.Errorf("count must be at least 1")
			}
			count = parsed
		}

		// Build and execute the git reset command
		gitCmd := fmt.Sprintf("git reset --soft HEAD~%d", count)
		
		if err := executor.Execute(gitCmd); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

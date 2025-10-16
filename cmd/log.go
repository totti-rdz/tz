package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/executor"
)

var (
	logAllFlag bool
)

var logCmd = &cobra.Command{
	Use:     "log [count]",
	Aliases: []string{"l"},
	Short:   "Show git log in compact format",
	Long: `Show git commit log in one-line format (git log --oneline).

By default shows all commits. Pass a count to limit the output.
Use -a flag to show full git log format instead of oneline.

Examples:
  tz log        # Show all commits (git log --oneline)
  tz l          # Same, using alias
  tz log 5      # Show last 5 commits (git log --oneline -5)
  tz l 10       # Show last 10 commits
  tz log -a     # Show full log (git log)
  tz log 5 -a   # Show last 5 commits in full format`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var gitCmd string

		if logAllFlag {
			// Full git log format
			gitCmd = "git log"
			
			// Add count limit if provided
			if len(args) > 0 {
				countStr := strings.TrimPrefix(args[0], "-")
				count, err := strconv.Atoi(countStr)
				if err != nil {
					return fmt.Errorf("invalid count: %s (must be a number)", args[0])
				}
				if count < 1 {
					return fmt.Errorf("count must be at least 1")
				}
				gitCmd = fmt.Sprintf("%s -%d", gitCmd, count)
			}
		} else {
			// Oneline format (default)
			gitCmd = "git log --oneline"
			
			// Add count limit if provided
			if len(args) > 0 {
				countStr := strings.TrimPrefix(args[0], "-")
				count, err := strconv.Atoi(countStr)
				if err != nil {
					return fmt.Errorf("invalid count: %s (must be a number)", args[0])
				}
				if count < 1 {
					return fmt.Errorf("count must be at least 1")
				}
				gitCmd = fmt.Sprintf("%s -%d", gitCmd, count)
			}
		}

		if err := executor.Execute(gitCmd); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().BoolVarP(&logAllFlag, "all", "a", false, "Show full git log instead of oneline format")
}

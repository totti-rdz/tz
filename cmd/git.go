package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/totti-rdz/tz/internal/executor"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Aliases: []string{"f"},
	Short:   "Git fetch",
	Long:    `Run git fetch to update remote tracking branches.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return executor.Execute("git fetch")
	},
}

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:     "branch <name>",
	Aliases: []string{"br"},
	Short:   "Create and checkout a new branch, or switch to previous with '-'",
	Long: `Create a new branch and switch to it (git checkout -b).
Use '-' as the branch name to switch to the previous branch.

Examples:
  tz branch feature-x    # Create and checkout new branch
  tz br -                # Switch to previous branch`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		branchName := args[0]

		// Special case: switch to previous branch
		if branchName == "-" {
			return executor.Execute("git checkout -")
		}

		// Normal case: create and checkout new branch
		return executor.Execute(fmt.Sprintf("git checkout -b %s", branchName))
	},
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"s"},
	Short:   "Git status",
	Long:    `Show the working tree status (git status).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return executor.Execute("git status")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(branchCmd)
	rootCmd.AddCommand(statusCmd)
}

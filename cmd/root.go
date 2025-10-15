package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tz",
	Short: "A CLI tool to streamline development commands",
	Long: `tz is a lightweight CLI tool that lets you use the same commands 
across all your projects, regardless of language or framework.

Map project-specific commands once, then use simple shortcuts everywhere:
  tz i    - Install dependencies
  tz d    - Start dev server
  tz t    - Run tests
  tz b    - Build project
  tz c    - Clear build artifacts

Universal git shortcuts work everywhere:
  tz f    - Git fetch
  tz s    - Git status
  tz br   - Create and checkout branch`,
	Version: "0.1.0",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// These variables are set during build time using -ldflags
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  "Print detailed version information about the woa CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("woa %s\n", version)
		fmt.Printf("  commit: %s\n", commit)
		fmt.Printf("  built:  %s\n", date)
	},
}

package cmd

import (
	"fmt"

	"github.com/20326/vega/app/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Show the version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", version.Version.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

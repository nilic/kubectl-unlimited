package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print kubectl-unlimited version",
	Long:  `Print kubectl-unlimited version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kubectl-unlimited v%s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

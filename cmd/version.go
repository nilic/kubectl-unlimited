package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

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

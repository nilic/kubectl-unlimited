package cmd

import (
	"github.com/nilic/kubectl-unlimited/unlimited"
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Display information about running containers with no memory limits set",
	Long:  `Display information about running containers with no memory limits set`,
	Run: func(cmd *cobra.Command, args []string) {
		config.SetCheckMemory()
		unlimited.ShowUnlimited(config)
	},
}

func init() {
	rootCmd.AddCommand(memoryCmd)
}

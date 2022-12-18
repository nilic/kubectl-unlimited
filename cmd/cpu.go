package cmd

import (
	"github.com/nilic/kubectl-unlimited/unlimited"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Display information about running containers with no CPU limits set",
	Long:  `Display information about running containers with no CPU limits set`,
	Run: func(cmd *cobra.Command, args []string) {
		unlimited.ShowUnlimited(kubeConfig, kubeContext, namespace, labels, true, false)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
}

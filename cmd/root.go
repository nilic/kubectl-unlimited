package cmd

import (
	"os"

	"github.com/nilic/kubectl-unlimited/unlimited"
	"github.com/spf13/cobra"
)

var (
	kubeConfig  string
	kubeContext string
	namespace   string
	labels      string

	rootCmd = &cobra.Command{
		Use:   "kubectl-unlimited",
		Short: "kubectl plugin for displaying information about running containers with no limits set.",
		Long:  "kubectl plugin for displaying information about running containers with no limits set.",

		Run: func(cmd *cobra.Command, args []string) {
			unlimited.ShowUnlimited(kubeConfig, kubeContext, namespace, labels, true, true)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&kubeConfig,
		"kubeconfig", "", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVarP(&kubeContext,
		"context", "", "", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().StringVarP(&namespace,
		"namespace", "n", "", "only analyze containers in this namespace (by default all containers from all namespaces are shown)")
	rootCmd.PersistentFlags().StringVarP(&labels,
		"labels", "l", "", "labels to filter pods with")
}

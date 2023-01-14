package cmd

import (
	"fmt"

	"github.com/nilic/kubectl-unlimited/unlimited"
	"github.com/spf13/cobra"
)

var (
	config = &unlimited.Config{}

	rootCmd = &cobra.Command{
		Use:   "kubectl-unlimited",
		Short: "kubectl plugin for displaying information about running containers with no limits set.",
		Long:  "kubectl plugin for displaying information about running containers with no limits set.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := config.Validate(); err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			config.SetCheckCPU()
			config.SetCheckMemory()
			unlimited.Show(config)
		},
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.KubeConfig,
		"kubeconfig", "", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVarP(&config.KubeContext,
		"context", "", "", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().StringVarP(&config.Namespace,
		"namespace", "n", "", "only analyze containers in this namespace")
	rootCmd.PersistentFlags().BoolVarP(&config.AllNamespaces,
		"all-namespaces", "A", false, "analyze containers in all namespaces")
	rootCmd.PersistentFlags().StringVarP(&config.Labels,
		"labels", "l", "", "labels to filter pods with")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFormat,
		"output", "o", "table",
		fmt.Sprintf("output format, one of: %v", unlimited.SupportedOutputFormats))
}

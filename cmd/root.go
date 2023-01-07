package cmd

import (
	"fmt"
	"log"

	"github.com/nilic/kubectl-unlimited/unlimited"
	"github.com/spf13/cobra"
)

var (
	config = &unlimited.Config{
		CheckCPU:    true,
		CheckMemory: true,
	}

	rootCmd = &cobra.Command{
		Use:   "kubectl-unlimited",
		Short: "kubectl plugin for displaying information about running containers with no limits set.",
		Long:  "kubectl plugin for displaying information about running containers with no limits set.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config.Validate()
		},
		Run: func(cmd *cobra.Command, args []string) {
			unlimited.ShowUnlimited(config)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config.KubeConfig,
		"kubeconfig", "", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVarP(&config.KubeContext,
		"context", "", "", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().StringVarP(&config.Namespace,
		"namespace", "n", "", "only analyze containers in this namespace")
	rootCmd.PersistentFlags().StringVarP(&config.Labels,
		"labels", "l", "", "labels to filter pods with")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFormat,
		"output", "o", "table",
		fmt.Sprintf("output format, one of: %v", unlimited.SupportedOutputFormats))
}

package cmd

import (
	"fmt"
	"log"

	"github.com/nilic/kubectl-unlimited/unlimited"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var (
	kubeConfig   string
	kubeContext  string
	namespace    string
	labels       string
	outputFormat string

	rootCmd = &cobra.Command{
		Use:   "kubectl-unlimited",
		Short: "kubectl plugin for displaying information about running containers with no limits set.",
		Long:  "kubectl plugin for displaying information about running containers with no limits set.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !slices.Contains(unlimited.SupportedOutputFormats, outputFormat) {
				log.Fatalf("error: invalid output format, please choose one of: %v\n", unlimited.SupportedOutputFormats)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			unlimited.ShowUnlimited(kubeConfig, kubeContext, namespace, labels, outputFormat, true, true)
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
	rootCmd.PersistentFlags().StringVarP(&kubeConfig,
		"kubeconfig", "", "", "path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVarP(&kubeContext,
		"context", "", "", "name of the kubeconfig context to use")
	rootCmd.PersistentFlags().StringVarP(&namespace,
		"namespace", "n", "", "only analyze containers in this namespace (by default all containers from all namespaces are shown)")
	rootCmd.PersistentFlags().StringVarP(&labels,
		"labels", "l", "", "labels to filter pods with")
	rootCmd.PersistentFlags().StringVarP(&outputFormat,
		"output", "o", "table",
		fmt.Sprintf("output format, one of: %v", unlimited.SupportedOutputFormats))
}

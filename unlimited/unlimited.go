package unlimited

import (
	"log"

	"golang.org/x/exp/slices"
)

type Config struct {
	KubeConfig   string
	KubeContext  string
	Namespace    string
	Labels       string
	OutputFormat string
	CheckCPU     bool
	CheckMemory  bool
}

func (c *Config) Validate() {
	if !slices.Contains(SupportedOutputFormats, c.OutputFormat) {
		log.Fatalf("error: invalid output format, please choose one of: %v\n", SupportedOutputFormats)
	}
}

func (c *Config) SetCheckCPU() {
	c.CheckCPU = true
}

func (c *Config) SetCheckMemory() {
	c.CheckMemory = true
}

func ShowUnlimited(c *Config) {
	clientset, err := getKubeClientset(c.KubeConfig, c.KubeContext)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	pods, err := getPods(clientset, c.Namespace, c.Labels)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	containerList := buildContainerList(pods, c.CheckCPU, c.CheckMemory)

	err = containerList.printContainers(c.OutputFormat)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

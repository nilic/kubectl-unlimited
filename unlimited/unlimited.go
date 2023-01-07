package unlimited

import (
	"fmt"
	"log"

	"golang.org/x/exp/slices"
	"k8s.io/client-go/tools/clientcmd"
)

type Config struct {
	KubeConfig    string
	KubeContext   string
	Namespace     string
	Labels        string
	OutputFormat  string
	AllNamespaces bool
	CheckCPU      bool
	CheckMemory   bool
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

func (c *Config) SetNamespace(clientconfig clientcmd.ClientConfig) error {
	if c.AllNamespaces {
		c.Namespace = ""
	} else if c.Namespace == "" {
		ctxNamespace, _, err := clientconfig.Namespace()
		if err != nil {
			return fmt.Errorf("error reading namespace from current context: %s", err.Error())
		}
		c.Namespace = ctxNamespace
	}

	return nil
}

func ShowUnlimited(c *Config) {
	clientconfig := getKubeConfig(c.KubeConfig, c.KubeContext)
	clientset, err := getKubeClientset(clientconfig)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	if err = c.SetNamespace(clientconfig); err != nil {
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

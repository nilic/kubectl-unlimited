package unlimited

import (
	"fmt"

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

func (c *Config) Validate() error {
	if !slices.Contains(SupportedOutputFormats, c.OutputFormat) {
		return fmt.Errorf("error: invalid output format, please choose one of: %v", SupportedOutputFormats)
	}

	return nil
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

package unlimited

import (
	"log"
)

func Show(c *Config) {
	clientconfig := newClientConfig(c.KubeConfig, c.KubeContext)
	clientset, err := newClientset(clientconfig)
	if err != nil {
		log.Fatalf("error: unable to generate clientset: %v\n", err)
	}

	if err = c.SetNamespace(clientconfig); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	pods, err := getPods(clientset, c.Namespace, c.Labels)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	containerList := buildContainerList(pods, c.CheckCPU, c.CheckMemory)

	if err = containerList.print(c.OutputFormat); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

package unlimited

import (
	"log"
)

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

	if err = containerList.printContainers(c.OutputFormat); err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

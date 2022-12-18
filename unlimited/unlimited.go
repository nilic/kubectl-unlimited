package unlimited

import (
	"log"
)

func ShowUnlimited(kubeConfig string, kubeContext string, namespace string, labels string, checkCPU bool, checkMemory bool) {
	clientset, err := getKubeClientset(kubeConfig, kubeContext)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	pods, err := getPods(clientset, namespace, labels)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	containerList := buildContainerList(pods, checkCPU, checkMemory)

	printContainerList(containerList)
}

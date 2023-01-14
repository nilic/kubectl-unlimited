package unlimited

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func newClientConfig(kubeConfig string, kubeContext string) clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if kubeConfig != "" {
		loadingRules.ExplicitPath = kubeConfig
	}

	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: kubeContext}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
}

func newClientset(c clientcmd.ClientConfig) (*kubernetes.Clientset, error) {
	config, err := c.ClientConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

package unlimited

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfig(kubeConfig string, kubeContext string) clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if kubeConfig != "" {
		loadingRules.ExplicitPath = kubeConfig
	}

	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: kubeContext}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
}

func getKubeClientset(c clientcmd.ClientConfig) (*kubernetes.Clientset, error) {
	config, err := c.ClientConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

// func getKubeClientset(kubeConfig string, kubeContext string) (*kubernetes.Clientset, error) {
// 	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
// 	if kubeConfig != "" {
// 		loadingRules.ExplicitPath = kubeConfig
// 	}

// 	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: kubeContext}
// 	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
// 		loadingRules,
// 		configOverrides,
// 	).ClientConfig()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// debug
// 	config2 := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
// 		loadingRules,
// 		configOverrides,
// 	)
// 	ns, _, _ := config2.Namespace()
// 	fmt.Printf("DEBUG: Current namespace is %s", ns)
// 	// end debug
// 	return kubernetes.NewForConfig(config)
// }

package unlimited

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type computeResources struct {
	CPURequest    resource.Quantity
	CPULimit      resource.Quantity
	memoryRequest resource.Quantity
	memoryLimit   resource.Quantity
}

type container struct {
	name      string
	podName   string
	namespace string
	resources computeResources
}

func getPods(clientset kubernetes.Interface, namespace string, labels string) (*corev1.PodList, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		return nil, fmt.Errorf("error listing Pods: %s", err.Error())
	}

	filteredPods := []corev1.Pod{}

	for _, p := range pods.Items {
		if p.Status.Phase != corev1.PodSucceeded && p.Status.Phase != corev1.PodFailed {
			filteredPods = append(filteredPods, p)
		}
	}

	pods.Items = filteredPods

	return pods, nil
}

func buildContainerList(pods *corev1.PodList, checkCPU bool, checkMemory bool) []container {
	containerList := []container{}

	for _, p := range pods.Items {
		for _, c := range p.Spec.Containers {
			if (checkCPU && c.Resources.Limits.Cpu().IsZero()) || (checkMemory && c.Resources.Limits.Memory().IsZero()) {
				containerList = append(containerList, container{
					name:      c.Name,
					podName:   p.Name,
					namespace: p.Namespace,
					resources: computeResources{
						CPURequest:    c.Resources.Requests["cpu"],
						CPULimit:      c.Resources.Limits["cpu"],
						memoryRequest: c.Resources.Requests["memory"],
						memoryLimit:   c.Resources.Limits["memory"],
					},
				})
			}
		}
	}

	return containerList
}

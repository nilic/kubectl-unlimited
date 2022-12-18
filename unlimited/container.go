package unlimited

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type computeResource struct {
	CPU    resource.Quantity `json:"cpu"`
	Memory resource.Quantity `json:"memory"`
}

type container struct {
	Name      string          `json:"name"`
	PodName   string          `json:"pod"`
	Namespace string          `json:"namespace"`
	Limits    computeResource `json:"limits"`
	Requests  computeResource `json:"requests"`
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
					Name:      c.Name,
					PodName:   p.Name,
					Namespace: p.Namespace,
					Limits: computeResource{
						CPU:    c.Resources.Limits["cpu"],
						Memory: c.Resources.Limits["memory"],
					},
					Requests: computeResource{
						CPU:    c.Resources.Requests["cpu"],
						Memory: c.Resources.Requests["memory"],
					},
				})
			}
		}
	}

	return containerList
}

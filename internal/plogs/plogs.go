package plogs

import (
	"bufio"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kha7iq/pl/internal/client"
	clr "github.com/kha7iq/pl/internal/colors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// logs retrieves and streams logs for the specified pod and container.
func logs(cs *client.Set, podName, namespace, highlightWord, containerName, label string, follow bool, tailLines int64, wg *sync.WaitGroup) error {
	defer wg.Done()

	// Define pod log options for streaming logs.
	podLogOpts := corev1.PodLogOptions{
		Follow:    follow,
		Container: containerName,
	}
	if tailLines > 0 {
		podLogOpts.TailLines = &tailLines
	}

	// Get logs stream for the specified pod and container.
	podLogs, err := cs.CoreV1Interface.Pods(namespace).GetLogs(podName, &podLogOpts).Stream(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get logs for %s: %v", podName, err)
	}
	defer podLogs.Close()

	reader := bufio.NewScanner(podLogs)
	for reader.Scan() {
		line := reader.Text()

		// Highlight word in the log line.
		if len(highlightWord) > 0 {
			line = clr.Highlight(line, highlightWord)
		}

		// Apply log level colors to the log line.
		line = clr.ColorizeLogLevels(line)

		// Print the log line with appropriate formatting.
		if len(label) > 0 {
			fmt.Printf("%v %v\n", clr.Cyan+podName+clr.Reset, line)
		} else {
			fmt.Println(line)
		}
	}
	time.Sleep(1 * time.Millisecond)

	return nil
}

// GetPodLogs fetches logs for the specified pod(s) based on label selector or single pod name.
func GetPodLogs(cs *client.Set, namespace, label, highlightWord, podName, containerName string, follow bool, tailLines int64, wg *sync.WaitGroup) error {
	podListOps := metav1.ListOptions{}
	if len(label) > 0 {
		podListOps.LabelSelector = label
	}

	// Fetch pods based on the provided namespace and label selector.
	pods, err := cs.Pods(namespace).List(context.Background(), podListOps)
	if err != nil {
		return fmt.Errorf("failed to get pods from cluster with labels: %v", err)
	}

	if len(label) > 0 {
		// Iterate over pods and stream logs for each one with a matching label.
		for _, pod := range pods.Items {
			containerName = pod.Spec.Containers[0].Name
			fmt.Printf("Pod: %s | Container: %v \n", clr.DimYellow+pod.Name+clr.Reset, clr.DimYellow+containerName+clr.Reset)
			wg.Add(1)
			go logs(cs, pod.Name, namespace, highlightWord, containerName, label, follow, tailLines, wg)
		}

	} else {
		// If label is not set, stream logs for the specified single pod.
		if len(containerName) <= 0 {
			containerName, err = getFirstContainerName(cs, namespace, podName)
			if err != nil {
				return fmt.Errorf("failed to get container name: %v", err)
			}
		}
		wg.Add(1)
		go logs(cs, podName, namespace, highlightWord, containerName, label, follow, tailLines, wg)
	}
	return nil
}

// getFirstContainerName retrieves the first container name from the specified pod.
func getFirstContainerName(cs *client.Set, namespace, podName string) (string, error) {
	pod, err := cs.Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get pod %s in namespace %s: %v", podName, namespace, err)
	}

	if len(pod.Spec.Containers) == 0 {
		return "", fmt.Errorf("pod %s in namespace %s has no containers", podName, namespace)
	}
	return pod.Spec.Containers[0].Name, nil
}

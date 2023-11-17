package client

import (
	"fmt"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Set provides the struct to interact with the Kubernetes API
type Set struct {
	client.Client
	corev1.CoreV1Interface
}

// New returns a *Set configured with the given kubeconfig
func New(kubeconfig string) (*Set, error) {
	// Get REST config from provided kubeconfig
	config, err := genericclioptions.NewConfigFlags(true).ToRESTConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create REST config: %v", err)
	}

	// Create a Set instance with the REST config
	clientSet := &Set{CoreV1Interface: corev1.NewForConfigOrDie(config)}

	return clientSet, nil
}

// GetCurrentNamespace retrieves the namespace from the current context.
// If the namespace is empty, it returns "default".
func (s *Set) GetCurrentNamespace() (string, error) {
	// Load kubeconfig for current context
	config, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		return "", fmt.Errorf("failed to load kubeconfig: %v", err)
	}

	currentContext := config.CurrentContext
	if currentContext == "" {
		return "", fmt.Errorf("current context not set in kubeconfig")
	}

	context := config.Contexts[currentContext]
	if context == nil {
		return "", fmt.Errorf("context %s not found in kubeconfig", currentContext)
	}

	namespace := context.Namespace
	if namespace == "" {
		return "default", nil // Set namespace to "default" if empty
	}

	return namespace, nil
}

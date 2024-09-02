package main

import (
	"context"
	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Load configuration
func getK8sConfig() *rest.Config {
	kubeconfig := filepath.Join("C:\\Users\\aoyang\\", ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func main() {
	cfg := getK8sConfig()

	scheme := runtime.NewScheme()
	chaosmeshv1alpha1.AddToScheme(scheme)

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme})
	if err != nil {
		panic(err.Error())
	}

	k8sClient.Create(context.Background(), &httpChaos)
}

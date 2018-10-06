package main

import (
	"fmt"
	"strings"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/kubernetes"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

func main() {
	// Get the kubernetes client
	client, err := kubernetes.GetClientset()
	if err != nil {
		logrus.Fatalf("kube client: %s", err)
	}

	// Call the list images function with the kubernetes client
	images, err := ListImages(client.CoreV1(), "")
	if err != nil {
		logrus.Fatalf("listing images: %s", err)
	}

	// Print the images contained in the cluster
	fmt.Println(strings.Join(images, "\n"))
}

// ListImages returns a list of container images running in the provided namespace
func ListImages(client v1.CoreV1Interface, namespace string) ([]string, error) {
	pl, err := client.Pods(namespace).List(meta_v1.ListOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "getting pods")
	}

	var images []string
	for _, p := range pl.Items {
		for _, c := range p.Spec.Containers {
			images = append(images, c.Image)
		}
	}

	return images, nil
}

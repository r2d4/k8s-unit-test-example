package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
)

func pod(namespace, image string) *v1.Pod {
	return &v1.Pod{ObjectMeta: meta_v1.ObjectMeta{Namespace: namespace}, Spec: v1.PodSpec{Containers: []v1.Container{{Image: image}}}}
}

func TestListImages(t *testing.T) {
	var tests = []struct {
		description string
		namespace   string
		expected    []string
		objs        []runtime.Object
	}{
		{"no pods", "", nil, nil},
		{"all namespaces", "", []string{"a", "b"}, []runtime.Object{pod("correct-namespace", "a"), pod("wrong-namespace", "b")}},
		{"filter namespace", "correct-namespace", []string{"a"}, []runtime.Object{pod("correct-namespace", "a"), pod("wrong-namespace", "b")}},
		{"wrong namespace", "correct-namespace", nil, []runtime.Object{pod("wrong-namespace", "b")}},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			client := fake.NewSimpleClientset(test.objs...)
			actual, err := ListImages(client.CoreV1(), test.namespace)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}
			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("%T differ (-got, +want): %s", test.expected, diff)
				return
			}
		})
	}
}

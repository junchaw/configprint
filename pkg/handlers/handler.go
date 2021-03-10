package handlers

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

type Handler interface {
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
}

type MockHandler struct {
}

func (s *MockHandler) ObjectCreated(obj interface{}) {
	pod, ok := obj.(*corev1.Pod)
	if ok {
		fmt.Printf("%s/%s created\n", pod.Namespace, pod.Name)
	}
}

func (s *MockHandler) ObjectDeleted(obj interface{}) {
	pod, ok := obj.(*corev1.Pod)
	if ok {
		fmt.Printf("%s/%s deleted\n", pod.Namespace, pod.Name)
	}
}

package main

import (
	"github.com/pkg/errors"
	"github.com/wbsnail/configprint/pkg/controller"
	"github.com/wbsnail/configprint/pkg/handlers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	klog.InitFlags(nil)

	klog.Info("reading config")

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		klog.Fatal(errors.Wrap(err, "read config error"))
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(errors.Wrap(err, "create clientset error"))
	}

	c, err := controller.New(clientset, &handlers.MockHandler{})
	if err != nil {
		klog.Fatal(errors.Wrap(err, "create controller error"))
	}

	stopCh := make(chan struct{})

	klog.Info("controller starting")

	c.Run(stopCh)

	klog.Info("controller started")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)

	<-sigterm

	klog.Info("terminated, exiting")
}

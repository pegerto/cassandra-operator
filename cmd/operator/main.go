package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/pegerto/cassandra-operator/pkg/controller"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
)



func main() {
	logrus.Info("Starting cassandra-operator")
	outside := flag.Bool("outside", false, "execute outside a cluster")

	var err error
	var config *rest.Config
	flag.Parse()

	if !*outside {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("",clientcmd.RecommendedHomeFile)
	}

	if err != nil {
		logrus.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	c := controller.New(clientset)
	c.Run()
}

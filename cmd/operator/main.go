package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/pegerto/cassandra-operator/pkg/controller"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	opkit "github.com/rook/operator-kit"
	"fmt"
	"os"
	"github.com/pegerto/cassandra-operator/pkg/cluster"
)



func main() {
	logrus.Info("Starting cassandra-operator")
	context, err := createContext()

	if err != nil {
		logrus.Errorf("failed to create context. %+v\n", err)
		os.Exit(1)
	}

	err = opkit.CreateCustomResources(*context, []opkit.CustomResource{cluster.ClusterResource})
	if err != nil {
		logrus.Errorf("failed to create custom resource. %+v\n", err)
		os.Exit(1)
	}

	c := controller.New(*context)
	c.Run()
}


func createContext() (*opkit.KubeContext, error) {
	// create the k8s client
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s client. %+v", err)
	}

	// initialize the sample resource
	httpCli, err := opkit.NewHTTPClient(cluster.ResourceGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to create http client. %+v", err)
	}

	return &opkit.KubeContext{
		MasterHost:  config.Host,
		Clientset:   clientset,
		KubeHTTPCli: httpCli.Client,
		RetryDelay:  6,
		MaxRetries:  15,
	}, nil
}
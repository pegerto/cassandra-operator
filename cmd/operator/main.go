package operator

import (
	"github.com/Sirupsen/logrus"
	"github.com/pegerto/cassandra-operator/pkg/controller"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main(){
	logrus.Info("Starting cassandra-operator")

	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	c := controller.New(clientset)
	c.Run()
}


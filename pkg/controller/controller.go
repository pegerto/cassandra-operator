package controller

import (
	"k8s.io/client-go/kubernetes"

	v1beta1extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	metav1 "k8s.io/client-go/pkg/api/v1"
	"github.com/Sirupsen/logrus"
)

type Controller struct {
	client *kubernetes.Clientset
}

func New(cli *kubernetes.Clientset) *Controller {
	return &Controller{
		client: cli,
	}
}

func (c *Controller) Run() {
	c.initResource()
}

func (c *Controller) initResource() {
	c.createTPR();

}

func (c *Controller) createTPR() error {

	tpr := &v1beta1extensions.ThirdPartyResource{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster.cassandra.com",
		},
		Versions: []v1beta1extensions.APIVersion{
			{Name: "v1beta1"},
		},
		Description: "Manage cassandra clusters",
	}

	_, err := c.client.ExtensionsV1beta1().ThirdPartyResources().Create(tpr)
	if err != nil  {
		logrus.Error("Error creating TPR:", err)
	}

	return err;
}

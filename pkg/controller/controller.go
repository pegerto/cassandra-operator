package controller

import (
	"k8s.io/client-go/kubernetes"
)

type Controller struct {
	client *kubernetes.Clientset
}

func New() *Controller {
	return &Controller{
	}
}

func (c *Controller) Run() {

}


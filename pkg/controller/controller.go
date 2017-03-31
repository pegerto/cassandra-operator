package controller

import "k8s.io/client-go/kubernetes"

type Controller struct {
	client *kubernetes.Clientset
}

func New(cli *kubernetes.Clientset) *Controller {
	return &Controller{
		client: cli,
	}
}


func (*Controller) Run() {

}


func (*Controller) initResource() {

}
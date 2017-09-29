package controller

import (
	"time"
	opkit "github.com/rook/operator-kit"
)

type Controller struct {
	context   opkit.KubeContext
}

func New(ctx opkit.KubeContext) *Controller {
	return &Controller{
		context: ctx,
	}
}

func (c *Controller) Run() {
	for {


		<-time.After(time.Second * time.Duration(c.context.RetryDelay))
	}
}


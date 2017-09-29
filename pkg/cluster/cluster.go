package cluster

import (
	opkit "github.com/rook/operator-kit"
)

const ResourceGroup = "cassandra.opencredo.com"


var ClusterResource = opkit.CustomResource{
	Name:        "cluster",
	Group:       ResourceGroup,
	Version:     opkit.V1Alpha1,
	Description: "A C* cluster running in k8s",
}

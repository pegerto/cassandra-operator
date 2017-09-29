package cluster

import (
	opkit "github.com/rook/operator-kit"
)

const ResourceGroup = "opencredo.com"


var ClusterResource = opkit.CustomResource{
	Name:        "cassandra",
	Group:       ResourceGroup,
	Version:     opkit.V1Alpha1,
	Description: "A cassandra cluster",
}

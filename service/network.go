package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetNetworkList() []types.NetworkResource {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	Networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}
	return Networks
}

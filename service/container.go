package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainerList() []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return containers
}

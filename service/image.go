package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetImageList() []types.ImageSummary {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{All: true})
	if err != nil {
		panic(err)
	}
	return images
}

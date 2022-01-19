package service

import (
	"context"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

func GetVolumeList() volume.VolumeListOKBody {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	filter := filters.NewArgs()

	Volumes, err := cli.VolumeList(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	return Volumes
}

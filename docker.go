package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

func docker() *cobra.Command {
	return &cobra.Command{
		Use: "docker",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
			if err != nil {
				panic(err)
			}

			containerListOptions := types.ContainerListOptions{All: true}
			containers, err := cli.ContainerList(ctx, containerListOptions)
			if err != nil {
				fmt.Println("Container list error")
				panic(err)
			}

			for _, container := range containers {
				if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
					fmt.Println("Container stop error")
					panic(err)
				}

				containerRemoveOptions := types.ContainerRemoveOptions{
					Force: true,
				}
				if err := cli.ContainerRemove(ctx, container.ID, containerRemoveOptions); err != nil {
					fmt.Println("Container remove error")
					panic(err)
				}
				fmt.Println(container.ID[:12])
			}

			images, err := cli.ImageList(ctx, types.ImageListOptions{})
			if err != nil {
				fmt.Println("Image list error")
				panic(err)
			}

			for _, image := range images {
				imageRemoveOptions := types.ImageRemoveOptions{Force: true, PruneChildren: true}
				_, err := cli.ImageRemove(ctx, image.ID, imageRemoveOptions)
				if err != nil {
					fmt.Println("Image remove error")
					panic(err)
				}
				fmt.Println("Deleted: ", image.ID)
			}

			networks, err := cli.NetworkList(ctx, types.NetworkListOptions{})
			if err != nil {
				fmt.Println("Network list error")
				panic(err)
			}

			for _, network := range networks {
				if !(network.Name == "none" || network.Name == "bridge" || network.Name == "host") {
					err := cli.NetworkRemove(ctx, network.ID)
					if err != nil {
						panic(err)
					}
					fmt.Println(network.ID)
				}
			}

			volumeArgs := filters.Args{}
			volumes, err := cli.VolumeList(ctx, volumeArgs)
			if err != nil {
				panic(err)
			}

			for _, volume := range volumes.Volumes {
				if err := cli.VolumeRemove(ctx, volume.Name, true); err != nil {
					fmt.Println("Volume remove error")
					panic(err)
				}
				fmt.Println(volume.Name)
			}

			return nil
		},
	}
}

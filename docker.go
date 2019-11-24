package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
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

			containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
			if err != nil {
				panic(err)
			}

			for _, container := range containers {
				if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
					panic(err)
				}
				fmt.Println(container.ID[:12])
			}

			images, err := cli.ImageList(ctx, types.ImageListOptions{})
			if err != nil {
				panic(err)
			}

			for _, image := range images {
				imageRemoveOptions := types.ImageRemoveOptions{Force: true, PruneChildren: true}
				id, err := cli.ImageRemove(ctx, image.ID, imageRemoveOptions)
				if err != nil {
					panic(err)
				}
				fmt.Println("Deleted: sha256:", id[:64])
			}

			return nil
		},
	}
}

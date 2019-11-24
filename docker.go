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
				fmt.Println("Stopping container: ", container.ID[:10], "...")
				if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
					panic(err)
				}
				fmt.Println("Successfully removed container: ", container.ID[:10], "...")
			}

			images, err := cli.ImageList(ctx, types.ImageListOptions{})
			if err != nil {
				panic(err)
			}

			for _, image := range images {
				fmt.Println("Stopping image: ", image.ID[:10], "...")

				imageRemoveOptions := types.ImageRemoveOptions{Force: true, PruneChildren: true}
				id, err := cli.ImageRemove(ctx, image.ID, imageRemoveOptions)
				if err != nil {
					panic(err)
				}
				fmt.Println("Successfully removed image ID: ", id[:10], "...")
			}

			return nil
		},
	}
}

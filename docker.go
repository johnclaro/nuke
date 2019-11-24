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

			return nil
		},
	}
}

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "nuke",
		Short:        "Remove all Docker related things",
		Long:         "Removes all Docker configs, containers, images, networks, nodes, plugins, secrets, services and volumes.",
		SilenceUsage: true,
	}

	cmd.AddCommand(docker())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "nuke",
		Short:        "Remove Docker stuff",
		SilenceUsage: true,
	}

	cmd.AddCommand(docker())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

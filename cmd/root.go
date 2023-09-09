package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/williamjhughes/flux/pkg/core"
)

var root = &cobra.Command{
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("expected 1 argument, got %d", len(args))
		}

		return nil
	},
	Use: "flux [script]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			core.UseFile(args[0])
			return
		}

		core.UseRepl()
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

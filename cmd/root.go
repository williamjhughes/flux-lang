package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "flux",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, Flux. 👋")
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

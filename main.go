package main

import (
	"log"

	"github.com/williamjhughes/flux/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}

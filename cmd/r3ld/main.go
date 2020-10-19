package main

import (
	"os"

	"github.com/relevant-community/r3l/cmd/r3ld/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}

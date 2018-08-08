package main

import (
	"os"

	"github.com/dtrv/kstatbeat/cmd"

	_ "github.com/dtrv/kstatbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

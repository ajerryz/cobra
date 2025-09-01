package main

import (
	"os"

	"github.com/ajerryz/cobra-apps/internal/daib/cmd"
)

func main() {

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}

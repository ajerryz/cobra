package main

import (
	"fmt"
	"os"

	"github.com/ajerryz/cobra-apps/internal/helloworld/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}

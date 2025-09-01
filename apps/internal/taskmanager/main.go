package taskmanager

import (
	"fmt"
	"os"

	"github.com/ajerryz/cobra-apps/internal/taskmanager/cmd"
)

func Main() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

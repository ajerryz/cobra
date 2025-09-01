package cmd

import "github.com/spf13/cobra"

var removeTaskCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a task",
}

func init() {
	rootCmd.AddCommand(removeTaskCmd)
}

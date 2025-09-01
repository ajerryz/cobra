package cmd

import "github.com/spf13/cobra"

var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "list current tasks",
}

func init() {
	rootCmd.AddCommand(removeTaskCmd)
}

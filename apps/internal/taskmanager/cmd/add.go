package cmd

import "github.com/spf13/cobra"

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long:  `add a task to you taskManager`,
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
}

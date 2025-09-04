package cmd

import "github.com/spf13/cobra"

var removeTaskCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove task",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(removeTaskCmd)
}

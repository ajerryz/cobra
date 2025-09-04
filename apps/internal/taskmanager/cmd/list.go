package cmd

import "github.com/spf13/cobra"

var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}

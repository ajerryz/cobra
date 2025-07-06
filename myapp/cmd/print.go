package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	printCmd.Flags().StringP("message", "m", "", "message to print")
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print message",
	Long:  `Print the message`,
	Run: func(cmd *cobra.Command, args []string) {
		flag := cmd.Flag("message")
		v := flag.Value.String()
		fmt.Println(v)
	},
}

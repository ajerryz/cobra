package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "myapp",
	Short:   "myapp a demo.",
	Version: "1.0.0",
	Long:    `myapp is a CLI tool to manage your myapp.`,
	Args:    cobra.MinimumNArgs(1),
}

// Execute 自定义执行入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(printCmd)
}

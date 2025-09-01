package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helloworld",
	Short: "this is short message",
	Long:  "这是一个cobra演示程序",
	Run: func(cmd *cobra.Command, args []string) {
		// 如果没有指定子命令，显示帮助信息
		fmt.Println(cmd.Name())
		fmt.Println(args)
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		config := cmd.Flag("config").Value.String()
		fmt.Println("欢迎来到 cobra 的演示程序")
		fmt.Printf("输入的config=%v\n", config)
	},
}

// Execute 执行rootCmd
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("config", "c", "", "config path")
}

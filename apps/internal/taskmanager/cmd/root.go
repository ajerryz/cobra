package cmd

import (
	"github.com/ajerryz/cobra-apps/internal/taskmanager/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskmanager [command]",                                                      //命令名称(在终端输入的命令)
	Short: "task manager",                                                               // 短描述，在help中显示
	Long:  " task-manager is a cobra project sample, you can add,list,remove your task", // 长描述
	Run: func(cmd *cobra.Command, args []string) { // 根命令的执行逻辑
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
	},
	Version: version.Version,
}

// Execute 程序的入口，调用根命令Execute方法。开始进行解析和执行
func Execute() error {
	return rootCmd.Execute()
}

// init 初始化函数(golang内容)，在这里可以定义全局标识
func init() {
	//
}

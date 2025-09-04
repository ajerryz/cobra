package cmd

import (
	"github.com/ajerryz/cobra-apps/internal/taskmanager/storage"
	"github.com/spf13/cobra"
)

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task",
	Long:  `add a task to you taskManager`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			cmd.Printf("args:%v\n", args)
		}

		index, _ := cmd.Flags().GetInt("index")
		name, _ := cmd.Flags().GetString("name")
		desc, _ := cmd.Flags().GetString("desc")

		err := storage.AppendTask(&storage.Task{
			Index: index,
			Name:  name,
			Desc:  desc,
		})
		if err != nil {
			cmd.Printf("append fail! error:%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	addTaskCmd.Flags().IntP("index", "i", 0, "index")
	addTaskCmd.Flags().StringP("name", "n", "taskName", "task name")
	addTaskCmd.Flags().StringP("desc", "d", "taskDesc", "task description")
	addTaskCmd.Flags().BoolP("force", "f", false, "force overwrite")

	_ = addTaskCmd.MarkFlagRequired("name")
}

package cmd

import (
	"fmt"
	"todolist/utils"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "快速查看todolist",
	Long:  `快速查看todolist，使用方式：todolist list`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todolist := listTodolist()
		if todolist == nil {
			fmt.Printf("todo列表还是空的，快去添加todo吧\n")
		}
		for _, todo := range todolist {
			fmt.Printf("%+v", todo)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTodolist() []*utils.TodolistIndex {
	indexes := utils.TodoListIndexes{Workdir: Workdir()}
	return indexes.List()
}

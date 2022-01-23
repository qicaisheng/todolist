package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add <title>",
	Short:   "快速添加todo",
	Long:    `快速添加todo，使用方式：todolist add "XXX"`,
	Example: `todolist add "TODO 1"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		todoId := todolist.AddTodo(title)
		fmt.Printf("\"%s\"创建成功，todoId: %v\n", title, todoId)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

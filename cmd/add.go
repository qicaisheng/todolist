package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
	"todolist/context"
)

var addCmd = &cobra.Command{
	Use:     "add <title>",
	Short:   "快速添加todo",
	Long:    `快速添加todo，使用方式：todolist add "XXX"`,
	Example: `todolist add "TODO 1"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		todoId := addTodo(title)
		fmt.Printf("\"%s\"创建成功，todoId: %v\n", title, todoId)
	},
}

func addTodo(title string) int {
	indexes := context.TodoListIndexes{Workdir: Workdir()}
	todoId := indexes.NewTodoId()
	fileName := fmt.Sprintf("%v-%s.md", strconv.Itoa(todoId), title)
	filePath := filepath.Join(viper.GetString("workdir"), fileName)
	todoItem := fmt.Sprintf("# %v-%s\n## status\n%s\n", strconv.Itoa(todoId), title, "OPEN")
	err := os.WriteFile(filePath, []byte(todoItem), 0644)
	if err != nil {
		_ = fmt.Errorf("write file error: %v", err)
	}
	indexes.AppendCreatedTodo(context.TodolistIndex{TodoId: todoId, Title: title, Status: "OPEN"})
	return todoId
}

func init() {
	rootCmd.AddCommand(addCmd)
}

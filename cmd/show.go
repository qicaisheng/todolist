package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"todolist/utils"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <todoId>",
	Short: "快速查看todo详情",
	Long:  `快速查看todo详情，使用方式：todolist show todoId`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("todoId is not valid number")
			os.Exit(1)
		}
		fmt.Println(getTodo(todoId))
	},
}

func getTodo(todoId int) string {
	indexes := utils.TodoListIndexes{Workdir: Workdir()}
	indexOf := indexes.IndexOf(todoId)
	fileName := fmt.Sprintf("%v-%s.md", todoId, indexOf.Title)

	todoItemFile := filepath.Join(Workdir(), fileName)
	todoDetail, err := os.ReadFile(todoItemFile)
	if err != nil {
		fmt.Printf("get todo detail error: %v\n", err)
		os.Exit(1)
	}
	return string(todoDetail)
}

func init() {
	rootCmd.AddCommand(showCmd)
}

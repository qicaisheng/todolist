package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
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
		fmt.Println(todolist.GetTodo(todoId))
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

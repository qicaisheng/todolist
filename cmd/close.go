package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close <todoId>",
	Short: "快速关闭todo",
	Long:  `快速关闭todo，使用方式：todolist close todoId`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("todoId is not valid number")
			os.Exit(1)
		}
		todolist.CloseTodo(todoId)
		fmt.Printf("todo %v 已经关闭\n", todoId)
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close <todoId>",
	Short: "快速关闭todo",
	Long:  `快速关闭todo，使用方式：todolist close todoId`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("close called, todoId is %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(closeCmd)
}

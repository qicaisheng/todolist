package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show <todoId>",
	Short: "快速查看todo详情",
	Long:  `快速查看todo详情，使用方式：todolist show todoId`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("show called, todoId is %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

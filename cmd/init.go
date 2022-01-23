package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "初始化todolist文件夹",
	Long:    `初始化todolist文件夹，使用方式：todolist init`,
	Example: `todolist init`,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		todolist.InitTodolist()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

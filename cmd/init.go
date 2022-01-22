package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"todolist/utils"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "初始化todolist文件夹",
	Long:    `初始化todolist文件夹，使用方式：todolist init`,
	Example: `todolist init`,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("init called")
		return initTodolist()
	},
}

func initTodolist() error {
	err := utils.InitTodolistIndexesFile(Workdir())
	if err != nil {
		return fmt.Errorf("init toolist error: %v", err)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}

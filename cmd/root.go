package cmd

import (
	"fmt"
	"os"
	"todolist/context"

	"github.com/spf13/cobra"
)

var todolist context.Todolist
var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "高效快捷进行todolist管理",
	Long:  `todolist作为一个CLI工具旨在高效快捷进行todolist管理，方面日常快速创建、跟进todo`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("get user home dir error: %v\n", err)
		os.Exit(1)
	}
	todolist = context.Todolist{Workdir: dirname}
	todolist.InitTodolist()
}

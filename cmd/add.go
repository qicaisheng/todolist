package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var addCmd = &cobra.Command{
	Use:     "add <title>",
	Short:   "快速添加todo",
	Long:    `快速添加todo，使用方式：todolist add "XXX"`,
	Example: `todolist add "TODO 1"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		addTodo(title)
		fmt.Printf("add called, titile is %v\n", args[0])
	},
}

func addTodo(title string) {
	fileName := title + ".md"
	filePath := filepath.Join(viper.GetString("workdir"), fileName)
	err := os.WriteFile(filePath, nil, 0644)
	if err != nil {
		_ = fmt.Errorf("write file error: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
}

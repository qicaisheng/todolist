package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:   "modify <todoId>",
	Short: "快速修改todo",
	Long:  `快速修改todo，使用方式：todolist modify todoId`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("todoId is not valid number")
			os.Exit(1)
		}

		index := todolist.IndexOf(todoId)

		prompt := promptui.Prompt{
			Label:     "标题",
			Validate:  func(input string) error { return nil },
			Default:   index.Title,
			AllowEdit: true,
		}

		updatedTitle, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		index.Title = updatedTitle
		todolist.ModifyTodo(index)
		fmt.Printf("todo %v 已更新成功，更新为如下：\n%s", index.TodoId, todolist.GetTodo(index.TodoId))
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}

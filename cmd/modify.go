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

		prompt := promptui.Prompt{
			Label:     "Title",
			Validate:  func(input string) error { return nil },
			Default:   "default todo title",
			AllowEdit: true,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
		fmt.Printf("modify called, todoId is %v\n", todoId)
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}

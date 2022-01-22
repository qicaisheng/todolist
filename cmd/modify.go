package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:   "modify <todoId>",
	Short: "快速修改todo",
	Long:  `快速修改todo，使用方式：todolist modify todoId`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("modify called, todoId is %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(modifyCmd)
}

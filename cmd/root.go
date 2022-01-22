package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

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
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(".todolist")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf(".todolist is broken, read error: %v", err)
	}
}

func Workdir() string {
	return viper.GetString("workdir")
}

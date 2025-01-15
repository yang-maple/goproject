package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "myoot", // Command的name
	Long:    "asdas",
	Short:   "简短描述", // Command 描述
	Version: "v1.1.1",
	//Long:  "dasdasd", // Command 描述
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exec 1111 mycobra command:", args)
	}, // Command 的执行函数
}

func Execute() {
	rootCmd.Execute() // 调用执行函数
}

// 初始化函数
func init() {
	// 添加参数标志
	// Bool 型的参数
	rootCmd.PersistentFlags().Bool("viper", true, "是否采用viper")
	// string 类型的参数
	rootCmd.PersistentFlags().StringP("auth", "a", "", "wu") // 全局标志
	rootCmd.Flags().StringP("sorce", "s", "", "sorce")       // 局部标志
}

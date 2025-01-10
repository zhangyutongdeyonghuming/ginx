package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 根命令
var rootCmd = &cobra.Command{
	Use:   "ginx",
	Short: "A CLI application for ginx",
	Long:  `ginx is a CLI tool that demonstrates global options, commands, and command-specific options.`,
}

// 添加全局选项
func init() {
	rootCmd.AddCommand(createCmd) // 注册 create 子命令
}

// Execute 初始化并运行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

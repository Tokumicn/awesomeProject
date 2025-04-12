package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd) // 字符串工具
	rootCmd.AddCommand(timeCmd) // 时间工具
}

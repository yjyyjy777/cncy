/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// activemqCmd represents the activemq command
var activemqCmd = &cobra.Command{
	Use:   "activemq",
	Short: "检测activemq",
	Long:  `通过向activemq测试队列发送消息，并接收消息进行状态检测`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("activemq called")
	},
}

func init() {
	rootCmd.AddCommand(activemqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activemqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activemqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/INVITATION-RPC-GATEWAY/cmd/serve"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "INVITATION-RPC-GATEWAY",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use serve to start a server")
	},
}

func execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Init() {
	serve.ServeCmd.Flags().StringVarP(&serve.Config, "config", "c", "", "Config in file://config.json")
	rootCmd.AddCommand(serve.ServeCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	execute()
}

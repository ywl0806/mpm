/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "excute",
	Aliases: []string{"e", "run"},
	Short:   "Execute project",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("execute called")
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}

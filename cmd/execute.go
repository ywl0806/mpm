/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ywl0806/my-pj-manager/pkg/execute"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "excute",
	Aliases: []string{"e", "run"},
	Short:   "Execute project",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		execute.ExecuteProjectByNames(args)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}

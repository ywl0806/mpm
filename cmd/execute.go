/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ywl0806/my-pj-manager/pkg/ask"
	"github.com/ywl0806/my-pj-manager/pkg/execute"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "excute",
	Aliases: []string{"e", "run"},
	Short:   "Execute project",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		projectNames, err := ask.SelectProjects()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		execute.ExecuteProjectByNames(projectNames)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)
}

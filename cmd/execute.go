/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ywl0806/mpm/pkg/ask"
	"github.com/ywl0806/mpm/pkg/execute"
)

// executeCmd represents the execute command
var executeCmd = &cobra.Command{
	Use:     "excute",
	Aliases: []string{"e", "run"},
	Short:   "Execute project",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		multi, err := cmd.Flags().GetBool("multi")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var projectNames []string
		if multi {
			projectNames, err = ask.SelectProjects()
		} else {
			projectName, _ := ask.PickProject()
			projectNames = append(projectNames, projectName)

		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		execute.ExecuteProjectByNames(projectNames)
	},
}

func init() {
	executeCmd.Flags().BoolP("multi", "m", false, "Execute multiple projects")
	rootCmd.AddCommand(executeCmd)

}

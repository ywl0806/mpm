/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ywl0806/mpm/pkg/ask"
	"github.com/ywl0806/mpm/pkg/db/project"
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

		oneWindow, err := cmd.Flags().GetBool("one-window")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		isPick, err := cmd.Flags().GetBool("pick")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var projectNames []string

		if multi { // if multi is true, ask user to select multiple projects
			projectNames, err = ask.SelectProjects()
		} else { // if multi is false, ask user to select one project
			projectName, _ := ask.PickProject()
			projectNames = append(projectNames, projectName)
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		pickDirs := make([][]string, len(projectNames))
		if isPick {
			for index, pn := range projectNames {
				pj, err := project.FindByName(pn)
				if err == nil {
					pdir, err := ask.PickDirs(pj)
					pickDirs[index] = pdir
					if err != nil {
						fmt.Println(err.Error())
						return
					}
				}
			}

			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if oneWindow { // if oneWindow is true, execute multiple projects in one window
			execute.ExecuteProjectAsOneWindow(projectNames, pickDirs)
			return
		}
		execute.ExecuteProjectByNames(projectNames, pickDirs)
	},
}

func init() {
	executeCmd.Flags().BoolP("multi", "m", false, "Execute multiple projects")
	executeCmd.Flags().BoolP("one-window", "o", false, "Execute multiple projects in one window")
	executeCmd.Flags().BoolP("pick", "p", false, "Pick a project")

	rootCmd.AddCommand(executeCmd)

}

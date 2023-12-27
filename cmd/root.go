package cmd

import (
	"fmt"
	"os"

	// "github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/ywl0806/my-pj-manager/pkg/db"
	"github.com/ywl0806/my-pj-manager/pkg/execute"
)

var rootCmd = &cobra.Command{
	Use:   "mpm",
	Short: "mpm - a simple CLI to managment local projects",
	Long:  `mpm <- Run recently used commands`,
	Run: func(cmd *cobra.Command, args []string) {

		projects, _ := db.List()
		db.SortProjectsByLastUsed(&projects)

		execute.ExecuteProject(projects[0])

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

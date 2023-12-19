package mpm

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/ywl0806/my-pj-manager/pkg/mpm"
)

var rootCmd = &cobra.Command{
	Use:   "mpm",
	Short: "mpm - a simple CLI to managment local projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		directories, err := mpm.GetGithubDirectories()
		if err != nil {
			return errors.New(" Error occurred")
		}
		for _, directory := range directories {
			fmt.Println("execute code", directory)
			cmd := exec.Command("code", directory)

			fmt.Println(cmd.Output())

		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

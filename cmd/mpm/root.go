package mpm

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mpm",
	Short: "mpm - a simple CLI to managment local projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		dirPath, err := os.Getwd()
		if err != nil {
			return errors.New(" Can not get directory path")
		}

		fmt.Println(dirPath)

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

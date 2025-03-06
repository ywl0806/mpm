package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ywl0806/mpm/pkg/execute"
)

var rootCmd = &cobra.Command{
	Use:     "mpm",
	Short:   "mpm - a simple CLI to managment local projects",
	Long:    `mpm <- Run recently used commands`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {

		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Println("mpm version: ", cmd.Version)
			os.Exit(0)
		}

		execute.ExecuteRecentProject()
	},
}

func Execute() {
	rootCmd.Flags().BoolP("version", "v", false, "Version")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

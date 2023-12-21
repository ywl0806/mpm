package cmd

import (
	"errors"
	"fmt"
	"os"

	// "github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/ywl0806/my-pj-manager/pkg/util"
)

var rootCmd = &cobra.Command{
	Use:   "mpm",
	Short: "mpm - a simple CLI to managment local projects",
	RunE: func(cmd *cobra.Command, args []string) error {

		// options := []string{"hoge", "gege"}
		// var answer string
		// prompt := &survey.Select{
		// 	Message: "선택해: ",
		// 	Options: options,
		// }

		// survey.AskOne(prompt, &answer)
		// fmt.Println(answer)
		// cmd.Help()
		directories, err := util.GetDirectories("", "", true)

		if err != nil {
			return errors.New(" Error occurred")
		}

		for _, directory := range directories {
			fmt.Println("execute code", directory)
			// cmd := exec.Command("code", directory)

			// fmt.Println(cmd.Output())

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

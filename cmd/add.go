/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ywl0806/mpm/pkg/add"
)

// addCmd represents the init command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		deep, _ := cmd.Flags().GetInt("deep")
		name, _ := cmd.Flags().GetString("name")
		isAll, _ := cmd.Flags().GetBool("all")
		add.Add(isAll, name, deep)

	},
}

func init() {
	addCmd.Flags().IntP("deep", "d", 0, "디렉토리의 계층")
	addCmd.Flags().StringP("name", "n", "", "프로젝트 이름")
	addCmd.Flags().BoolP("all", "a", false, "현재 디렉토리 전부를 대상")
	rootCmd.AddCommand(addCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

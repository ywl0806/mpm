package mpm

import (
	"errors"
	"fmt"

	"os"
)

func GetGithubDirectories() ([]string, error) {
	currentDirectory, err := os.Getwd()

	if err != nil {
		return nil, errors.New(" Can not get directory path")
	}

	files, _ := os.ReadDir(currentDirectory)

	var githubDirectories []string
	for _, file := range files {
		if !file.Type().IsDir() {
			continue
		}
		if !CheckFileIsExist(".git", file.Name()) {
			githubDirectories = append(githubDirectories, file.Name())
		}
	}

	return githubDirectories, nil
}
func CheckFileIsExist(filename string, directory string) bool {
	files, _ := os.ReadDir(directory)

	for _, file := range files {
		if file.Name() == filename {
			fmt.Println()
			return true
		}
	}
	return false
}

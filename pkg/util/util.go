package util

import (
	"errors"
	"fmt"

	"os"
)

func GetDirectories(directory string, searchFileName string) ([]string, error) {

	var currentDirectory string

	if directory == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, errors.New(" Can not get directory path")
		}
		currentDirectory = dir
	} else {
		currentDirectory = directory
	}

	files, _ := os.ReadDir(currentDirectory)

	var githubDirectories []string
	for _, file := range files {
		if !file.Type().IsDir() {
			continue
		}
		if searchFileName == "" {
			githubDirectories = append(githubDirectories, file.Name())
			continue
		}
		if CheckFileIsExist(searchFileName, file.Name()) {
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

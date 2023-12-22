package util

import (
	"errors"
	"fmt"
	"io/fs"

	"os"
)

// 디렉토리안의 디렉토리의 리스트를 전달
//
//   - `directory` == "" 일경우 현재의 디렉토리를 탐색
//   - `searchFileName` -> 필터링대상 파일
//   - `include` -> 필터링 대상 포함 여부 ture이면 필터링 대상 파일이 포함된 디렉토리만 반환
func GetDirectories(directory string, searchFileName string, include bool) ([]fs.DirEntry, error) {

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

	var directories []fs.DirEntry
	for _, file := range files {
		if !file.Type().IsDir() {
			continue
		}
		if searchFileName == "" {
			directories = append(directories, file)
			continue
		}
		if CheckFileIsExist(searchFileName, file.Name()) == include {
			directories = append(directories, file)
		}
	}

	return directories, nil
}

// 디렉토리안에 특정 파일이 존재하는지 검사
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

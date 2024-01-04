package util

import (
	"errors"
	"path/filepath"

	"os"
)

// 디렉토리안의 디렉토리의 리스트를 전달
//
//   - `directory` == "" 일경우 현재의 디렉토리를 탐색
//   - `searchFileName` -> 필터링대상 파일
//   - `include` -> 필터링 대상 포함 여부 ture이면 필터링 대상 파일이 포함된 디렉토리만 반환
func GetDirectories(directory string, searchFileName string, include bool, deep int) ([]string, error) {

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

	var directories []string
	findDirectory(currentDirectory, "", searchFileName, include, &directories, deep)

	return directories, nil
}

// 재귀적으로 디렉토리 검사
func findDirectory(root string, dir string, searchFileName string, include bool, directories *[]string, deep int) {
	files, _ := os.ReadDir(filepath.Join(root, dir))

	for _, file := range files {
		nextDir := filepath.Join(dir, file.Name())
		if !file.Type().IsDir() || file.Name()[0:1] == "." {
			continue
		}

		if searchFileName == "" || CheckFileIsExist(searchFileName, nextDir) == include {
			*directories = append(*directories, filepath.Join(dir, file.Name()))
		}

		if deep != 0 {

			findDirectory(root, nextDir, searchFileName, include, directories, deep-1)
		}
	}
}

// 디렉토리안에 특정 파일이 존재하는지 검사
func CheckFileIsExist(filename string, directory string) bool {
	files, _ := os.ReadDir(directory)

	for _, file := range files {
		if file.Name() == filename {
			return true
		}
	}
	return false
}

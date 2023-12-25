package register

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ywl0806/my-pj-manager/pkg/db"
	"github.com/ywl0806/my-pj-manager/pkg/util"
)

// 현재 디렉토리를 등록
func Add(isAll bool, name string) {
	var newProject db.Project = db.Project{}

	// 이름 지정이 없을 경우 현재 디렉토리 이름으로함
	if name != "" {
		newProject.Name = name
	} else {
		newProject.Name = getCurrentDirectoryName()
	}

	var directoryNames []string
	// 현재 폴더의 모든 디렉토리를 포함시킴
	if !isAll {
		directoryNames = SurveyChooseDirectory()
	} else {
		directoriesFile, err := util.GetDirectories("", "", true)
		if err != nil {
			log.Println(err.Error())
			return
		}
		for _, dir := range directoriesFile {
			directoryNames = append(directoryNames, dir.Name())
		}
	}

	for _, name := range directoryNames {
		newProject.Directories = append(newProject.Directories, db.Directory{Path: name})
	}
	_, err := db.Add(newProject)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	fmt.Println("Success to Add")
}

func getCurrentDirectoryName() string {
	currentDirectory, _ := os.Getwd()
	currentDirectorySplit := strings.Split(currentDirectory, "/")
	currentDirectoryName := currentDirectorySplit[len(currentDirectorySplit)-1]
	return currentDirectoryName
}

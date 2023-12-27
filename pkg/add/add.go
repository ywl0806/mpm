package add

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ywl0806/my-pj-manager/pkg/ask"
	"github.com/ywl0806/my-pj-manager/pkg/db"
	"github.com/ywl0806/my-pj-manager/pkg/util"
)

// 현재 디렉토리를 등록
func Add(isAll bool, name string) {
	var newProject db.Project = db.Project{}
	currentPath, currentDirectoryName := getCurrentDirectory()

	newProject.Path = currentPath
	// 이름 지정이 없을 경우 설문
	if name != "" {
		newProject.Name = name
	} else {
		var err error
		newProject.Name, err = ask.InputName(currentDirectoryName)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	// 이름 중복체크
	if db.IsProjectNameDuplicate(newProject.Name) {
		fmt.Println("[" + newProject.Name + "]" + " name is duplicated")
		return
	}

	var directoryNames []string
	// 현재 폴더의 모든 디렉토리를 포함시킴
	if !isAll {
		var err error
		directoryNames, err = ask.SurveyChooseDirectory()
		if err != nil {
			log.Println(err.Error())
			return
		}
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

	commands, _ := ask.InputCmdEachDirectories(directoryNames)

	for _, name := range directoryNames {
		newProject.Directories = append(
			newProject.Directories,
			db.Directory{Path: name, Cmd: commands[name].Cmd, Options: commands[name].Options})
	}

	// add timestamp
	now := time.Now().String()
	newProject.Created_at = now
	newProject.Last_use_at = now

	// create to db
	_, dbAddErr := db.Add(newProject)

	if dbAddErr != nil {
		log.Fatalln(dbAddErr.Error())
		return
	}

	fmt.Println("Success to Add")
	fmt.Println(newProject)
}

func getCurrentDirectory() (string, string) {
	currentPath, _ := os.Getwd()
	currentDirectorySplit := strings.Split(currentPath, "/")
	currentDirectoryName := currentDirectorySplit[len(currentDirectorySplit)-1]
	return currentPath, currentDirectoryName
}

package add

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ywl0806/mpm/pkg/ask"
	"github.com/ywl0806/mpm/pkg/db/project"
	"github.com/ywl0806/mpm/pkg/util"
)

// 현재 디렉토리를 등록
func Add(isAll bool, name string, deep int) {
	var err error

	newProject := project.Project{}
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

	var directoryNames []string
	// 현재 폴더의 모든 디렉토리를 포함시킴
	if !isAll {
		directoryNames, err = ask.SurveyChooseDirectory(deep)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		directoryNames, err = util.GetDirectories("", "", true, deep)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	defaultCmd, err := ask.InputCmd("Default Command:")
	if err != nil {
		log.Println(err.Error())
		return
	}
	newProject.DefaultCmd = defaultCmd

	commands, _ := ask.InputCmdEachDirectories(directoryNames, defaultCmd)

	for _, name := range directoryNames {
		newProject.Directories = append(
			newProject.Directories,
			project.Directory{Path: name, Cmd: commands[name].Cmd, Options: commands[name].Options})
	}

	// add timestamp
	now := time.Now().String()
	newProject.Created_at = now
	newProject.Last_use_at = now

	// create to db
	err = project.Add(newProject)
	if err != nil {
		log.Fatalln(err.Error())
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

package register

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ywl0806/my-pj-manager/pkg/util"
)

// direcoty선택 질의
func SurveyChooseDirectory() []string {
	// git이 들어있는 디렉토리를 위로 정렬
	directoriesGit, _ := util.GetDirectories("", ".git", true)
	directoriesNoGit, _ := util.GetDirectories("", ".git", false)

	directories := append(directoriesGit, directoriesNoGit...)

	var directoryNames []string

	for _, dir := range directories {
		directoryNames = append(directoryNames, dir.Name())
	}
	var directoryNamseGit []string

	for _, dir := range directoriesGit {
		directoryNamseGit = append(directoryNamseGit, dir.Name())
	}
	prompt := &survey.MultiSelect{
		Message: "Choose directories",
		Options: directoryNames,
		Default: directoryNamseGit,
	}

	var answer []string
	err := survey.AskOne(prompt, &answer)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return answer
}

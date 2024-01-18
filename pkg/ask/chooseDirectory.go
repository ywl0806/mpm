package ask

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ywl0806/mpm/pkg/util"
)

// direcoty선택 질의
func SurveyChooseDirectory(deep int) ([]string, error) {

	directoriesGit, _ := util.GetDirectories("", ".git", true, deep)

	directories, _ := util.GetDirectories("", "", false, deep)

	prompt := &survey.MultiSelect{
		Message:  "Choose directories",
		Options:  append([]string{"."}, directories...),
		Default:  directoriesGit,
		PageSize: 20,
	}

	var answer []string
	err := survey.AskOne(prompt, &answer)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return answer, nil
}

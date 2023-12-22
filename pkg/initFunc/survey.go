package initFunc

import (
	"fmt"

	"github.com/ywl0806/my-pj-manager/pkg/util"
)

// direcoty선택 질의
func SurveyChooseDirectory() {
	// git이 들어있는 디렉토리를 위로 정렬
	directoriesGit, _ := util.GetDirectories("", ".git", true)
	directoriesNoGit, _ := util.GetDirectories("", ".git", false)

	directories := append(directoriesGit, directoriesNoGit...)

	fmt.Println(directories)

	var diretoryNames []string

	for _, dir := range directories {
		diretoryNames = append(diretoryNames, dir.Name())
	}

}

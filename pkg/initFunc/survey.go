package initfunc

import (
	"fmt"

	"github.com/ywl0806/my-pj-manager/pkg/util"
)

// direcoty선택 질의
func SurveyChooseDirectory() {
	directoriesGit, _ := util.GetDirectories("", ".git", true)
	directoriesNoGit, _ := util.GetDirectories("", ".git", false)

	directories := append(directoriesGit, directoriesNoGit...)

	fmt.Println(directories)
}

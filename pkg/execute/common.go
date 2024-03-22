package execute

import (
	"errors"
	"log"
	"os/exec"
	"time"

	"github.com/ywl0806/mpm/pkg/db/project"
	"github.com/ywl0806/mpm/pkg/util"
)

func executeProject(pj project.Project, pickDirs []string) {

	if pickDirs == nil {
		pickDirs = make([]string, 0)
	}

	pickDirSet := util.Array2Set(pickDirs)

	for _, dir := range pj.Directories {
		if len(pickDirs) != 0 && !pickDirSet[dir.Path] {
			continue
		}

		executer := exec.Command(dir.Cmd, pj.Path+"/"+dir.Path, dir.Options)

		if errors.Is(executer.Err, exec.ErrDot) {
			executer.Err = nil
		}
		if err := executer.Run(); err != nil {
			log.Fatal(err)
		}

	}

	updateTimestamp(pj)
}

func updateTimestamp(pj project.Project) {
	lastUseAt := time.Now().String()
	usage := pj.Usage + 1
	project.Update(pj.Name, &project.UpdateProject{
		Last_use_at: &lastUseAt,
		Usage:       &usage,
	})
}

func executeProjectAsOneWindow(pjs []project.Project, pickDirs [][]string) {
	paths := make([]string, 0)

	for index, pj := range pjs {
		for _, dir := range pj.Directories {
			pDirs := pickDirs[index]
			if len(pDirs) != 0 {
				if !util.Contains(pickDirs[index], dir.Path) {
					continue
				}
			}
			paths = append(paths, pj.Path+"/"+dir.Path)
		}
	}
	defaultCmd := pjs[0].DefaultCmd
	if defaultCmd == "" {
		defaultCmd = "code"
	}
	executer := exec.Command(defaultCmd, paths...)

	if errors.Is(executer.Err, exec.ErrDot) {
		executer.Err = nil
	}
	if err := executer.Run(); err != nil {
		log.Fatal(err)
	}

	for _, pj := range pjs {
		updateTimestamp(pj)
	}

}

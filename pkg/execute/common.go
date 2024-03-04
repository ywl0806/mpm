package execute

import (
	"errors"
	"log"
	"os/exec"
	"time"

	"github.com/ywl0806/mpm/pkg/db/project"
)

func executeProject(pj project.Project) {

	for _, dir := range pj.Directories {

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

func executeProjectAsOneWindow(pjs []project.Project) {
	paths := make([]string, len(pjs))

	for i, pj := range pjs {
		for _, dir := range pj.Directories {
			paths[i] = pj.Path + "/" + dir.Path
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

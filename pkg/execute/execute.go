package execute

import (
	"errors"
	"log"
	"os/exec"

	"github.com/ywl0806/my-pj-manager/pkg/db"
)

func ExecuteProject(project db.Project) {

	for _, dir := range project.Directories {

		executer := exec.Command(dir.Cmd, project.Path+"/"+dir.Path, dir.Options)

		if errors.Is(executer.Err, exec.ErrDot) {
			executer.Err = nil
		}
		if err := executer.Run(); err != nil {
			log.Fatal(err)
		}

	}
}

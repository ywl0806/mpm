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

	lastUseAt := time.Now().String()
	usage := pj.Usage + 1
	project.Update(pj.Name, &project.UpdateProject{
		Last_use_at: &lastUseAt,
		Usage:       &usage,
	})

}

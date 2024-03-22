package execute

import (
	"fmt"
	"log"
	"sort"

	"github.com/ywl0806/mpm/pkg/db/project"
)

func ExecuteRecentProject() {
	projects, _ := project.List()
	sort.Slice(projects, func(i, j int) bool { return projects[i].Last_use_at > projects[j].Last_use_at })
	if len(projects) == 0 {
		fmt.Println("No projects exist")
		return
	}
	executeProject(projects[0], nil)
}
func ExecuteProjectByNames(names []string, pickDirs [][]string) {

	projects := make([]project.Project, 0)

	for _, name := range names {
		pj, err := project.FindByName(name)
		if err != nil {
			log.Fatalln(err)
		}
		projects = append(projects, pj)
	}

	for index, pj := range projects {

		executeProject(pj, pickDirs[index])
	}

}

func ExecuteProjectAsOneWindow(names []string, pickDirs [][]string) {

	var projects []project.Project

	for _, name := range names {
		pj, err := project.FindByName(name)
		if err != nil {
			log.Fatalln(err)
		}
		projects = append(projects, pj)
	}

	executeProjectAsOneWindow(projects, pickDirs)
}

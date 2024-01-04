package execute

import (
	"log"
	"sort"

	"github.com/ywl0806/my-pj-manager/pkg/db/project"
)

func ExecuteRecentProject() {
	projects, _ := project.List()
	sort.Slice(projects, func(i, j int) bool { return projects[i].Last_use_at > projects[j].Last_use_at })
	executeProject(projects[0])
}
func ExecuteProjectByNames(names []string) {

	var projects []project.Project

	for _, name := range names {
		pj, err := project.FindByName(name)
		if err != nil {
			log.Fatalln(err)
		}
		projects = append(projects, pj)
	}

	for _, pj := range projects {
		executeProject(pj)
	}

}

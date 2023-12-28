package execute

import (
	"github.com/ywl0806/my-pj-manager/pkg/db"
)

func ExecuteRecentProject() {
	projects, _ := db.List()
	db.SortProjectsByLastUsed(&projects)
	executeProject(projects[0])
}
func ExecuteProjectByNames(names []string) {
	allProjects, _ := db.List()
	var projects []db.Project

	nameSet := map[string]bool{}

	for _, name := range names {
		nameSet[name] = true
	}

	for _, pj := range allProjects {
		if nameSet[pj.Name] {
			projects = append(projects, pj)
		}
	}

}

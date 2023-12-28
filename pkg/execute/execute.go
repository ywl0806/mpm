package execute

import (
	"sort"

	"github.com/ywl0806/my-pj-manager/pkg/db/project"
)

func ExecuteRecentProject() {
	projects, _ := project.List()
	sort.Slice(projects, func(i, j int) bool { return projects[i].Last_use_at > projects[j].Last_use_at })
	executeProject(projects[0])
}
func ExecuteProjectByNames(names []string) {
	allProjects, _ := project.List()
	var projects []project.Project

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

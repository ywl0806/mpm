package ask

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ywl0806/my-pj-manager/pkg/db/project"
)

func SelectProjects() (selectedPjs []string, err error) {
	var projects []project.Project
	projects, err = project.List()

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	var allPjNames []string
	sort.Slice(projects, func(i, j int) bool { return projects[i].Last_use_at > projects[j].Last_use_at })

	for _, pj := range projects {
		allPjNames = append(allPjNames, pj.Name)
	}

	prompt := &survey.MultiSelect{
		Message: "Choose projects",
		Options: allPjNames,
	}

	err = survey.AskOne(prompt, &selectedPjs)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return selectedPjs, err
}

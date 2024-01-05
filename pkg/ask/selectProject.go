package ask

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ywl0806/my-pj-manager/pkg/db/project"
)

func SelectProjects() (selectedPjs []string, err error) {
	var allPjNames []string

	allPjNames, err = getProjectNamesSortedByLastUseAt()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	prompt := &survey.MultiSelect{
		Message:  "Choose projects",
		Options:  allPjNames,
		PageSize: 20,
	}

	err = survey.AskOne(prompt, &selectedPjs)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return selectedPjs, err
}

func PickProject() (projectName string, err error) {
	var allPjNames []string

	allPjNames, err = getProjectNamesSortedByLastUseAt()

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	prompt := &survey.Select{
		Message:  "Pick project",
		Options:  allPjNames,
		PageSize: 20,
	}

	err = survey.AskOne(prompt, &projectName)

	return projectName, err
}

func getProjectNamesSortedByLastUseAt() ([]string, error) {
	projects, err := project.List()

	sort.Slice(projects, func(i, j int) bool { return projects[i].Last_use_at > projects[j].Last_use_at })

	var names []string
	for _, pj := range projects {
		names = append(names, pj.Name)
	}
	return names, err
}

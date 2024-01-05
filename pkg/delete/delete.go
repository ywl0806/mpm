package delete

import (
	"fmt"

	"github.com/ywl0806/my-pj-manager/pkg/ask"
	"github.com/ywl0806/my-pj-manager/pkg/db/project"
	"github.com/ywl0806/my-pj-manager/pkg/list"
)

func Delete() {
	target, err := ask.PickProject()

	if err != nil {
		return
	}
	pj, err := project.FindByName(target)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	list.ShowListDetail([]project.Project{pj})

	if !ask.Confirm("Delete this project \n") {
		return
	}

	err = project.Delete(target)

	if err != nil {
		fmt.Println("Can not delete the project")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Success to delete")

}

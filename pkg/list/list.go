package list

import (
	"fmt"
	"sort"

	"github.com/ywl0806/my-pj-manager/pkg/db"
)

func ShowList(detail bool) {
	list, err := db.List()
	sort.Slice(list, func(i, j int) bool { return list[i].Last_use_at > list[j].Last_use_at })
	if err != nil {
		fmt.Println("Show list error")
		return
	}

	if detail {
		showListDetail(list)
		return
	}
	showListOnlyName(list)

}

func showListOnlyName(projects []db.Project) {
	for _, pj := range projects {
		fmt.Println(pj.Name)
	}
}

func showListDetail(projects []db.Project) {
	for _, pj := range projects {
		fmt.Print("\n\n--------------------------------------\n\n")
		fmt.Printf("Name : %s \n", pj.Name)
		fmt.Printf("Path : %s \n", pj.Path)
		fmt.Printf("Usage : %d \n", pj.Usage)
		fmt.Println("Directories : ")
		for _, dir := range pj.Directories {
			fmt.Printf("\t[%s]: \n", dir.Path)
			fmt.Printf("\t  Path: %s \n", dir.Path)
			fmt.Printf("\t  Command: %s \n", dir.Cmd)
			fmt.Printf("\t  Options: %s \n", dir.Options)
		}
	}
}

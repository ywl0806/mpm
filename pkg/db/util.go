package db

import "sort"

func SortProjectsByLastUsed(list *[]Project) {
	project := *list

	sort.Slice(project, func(i, j int) bool { return project[i].Last_use_at > project[j].Last_use_at })
}

package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func saveProjects(projects []Project) error {
	data := getData()
	data.Projects = projects
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err.Error())
		return errors.New(err.Error())
	}
	err = os.WriteFile(DB_PATH, jsonData, 0644)
	if err != nil {
		log.Fatalln(err.Error())
		return errors.New(err.Error())
	}
	return nil
}

func getProjects() ([]Project, error) {
	rawFile, err := os.ReadFile(DB_PATH)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, errors.New(err.Error())
	}

	var data JsonData

	json.Unmarshal(rawFile, &data)

	return data.Projects, nil
}

func checkDuplicate(projects []Project, name string) bool {
	for _, project := range projects {
		if project.Name == name {
			return true
		}
	}
	return false
}

func getData() JsonData {
	var data JsonData
	jsonData, err := os.ReadFile(DB_PATH)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return data
}

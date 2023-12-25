package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var DB_PATH string = "tmp/data.json"

type JsonData struct {
	Projects []Project `json:"projects"`
}
type Project struct {
	Name        string      `json:"name"`
	Path        string      `json:"path"`
	Cmd         string      `json:"cmd"`
	Usage       int         `json:"usage"`
	Last_use_at string      `json:"last_use_at"`
	Directories []Directory `json:"directories"`
}

type Directory struct {
	Path    string `json:"path"`
	Cmd     string `json:"cmd"`
	Options string `json:"options"`
}

func init() {
	data, err := os.Open(DB_PATH)

	if err != nil && data == nil {
		log.Println("***init db***")

		projects := []Project{}
		initData := JsonData{Projects: projects}

		newJson, createErr := os.Create(DB_PATH)

		if createErr != nil {
			fmt.Println("json create error")
			return
		}
		defer newJson.Close()

		encoder := json.NewEncoder(newJson)

		err := encoder.Encode(initData)
		if err != nil {
			fmt.Println("JSON 쓰기 오류:", err)
			return
		}

		fmt.Println("data initalize")

	}

}
func List() ([]Project, error) {
	projects, err := getProjects()

	if err != nil {
		fmt.Println("Can not get data")
		return nil, err
	}

	return projects, nil
}
func Add(project Project) (Project, error) {
	projects, err := getProjects()

	if err != nil {
		fmt.Println("Can not get data")
		return Project{}, err
	}

	if checkDuplicate(projects, project.Name) {
		return Project{}, errors.New("duplicate project name")
	}

	projects = append(projects, project)
	saveProjects(projects)
	return project, nil

}

func Delete() {}

func Update() {}

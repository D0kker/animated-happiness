package sonarqube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type objGetAllModules struct {
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	} `json:"paging"`
	Components []Components
}

type Components struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Qualifier string `json:"qualifier"`
	Project   string `json:"project"`
}

type ProjectList struct {
	Projects []Components
}

type ProjectListModules struct {
	Back  ProjectList
	Front ProjectList
}

func (box *ProjectList) AddItem(item Components) []Components {
	box.Projects = append(box.Projects, item)
	return box.Projects
}

func GetAllModules(w http.ResponseWriter, r *http.Request) ProjectListModules {
	var bird objGetAllModules
	items := []Components{}
	ObjFront := ProjectList{items}
	ObjBack := ProjectList{items}
	i := 1

	for {
		route := fmt.Sprintf("http://devops/sonar/api/components/search?ps=500&qualifiers=TRK&p=%v", i)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", route, nil)
		req.Header.Set("Authorization", "Basic NTY5NGVhN2JmMDA1ZDFiYjM4ZTk4ZjkyMzRmOGI4MGMwODg1MzE0NDo=")
		res, _ := client.Do(req)
		responseData, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal([]byte(responseData), &bird)
		for _, element := range bird.Components {
			if strings.Contains(element.Key, "bel:") {
				item1 := Components{Key: element.Key, Name: element.Name, Qualifier: element.Qualifier, Project: element.Project}
				if strings.Contains(element.Key, ":lrf:") {
					ObjFront.AddItem(item1)
				}
				if strings.Contains(element.Key, ":lrj:") {
					ObjBack.AddItem(item1)
				}
			}
		}

		if (bird.Paging.PageIndex * bird.Paging.PageSize) > bird.Paging.Total {
			break
		} else {
			i = i + 1
		} //fmt.Println(bird.Paging)
	}

	fmt.Println(len(ObjFront.Projects))
	fmt.Println(len(ObjBack.Projects))

	return ProjectListModules{Back: ObjBack, Front: ObjFront}
}
package main

import (
	"cualquier_vaina/sonarqube"
	"cualquier_vaina/html"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
  "os"
  "path/filepath"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func main() {
	port := 1515

	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/getAllModules", getAllModules)
	http.HandleFunc("/metricsDaily", getAllModules)
	http.HandleFunc("/metricsMonthly", getAllModules)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func getAllModules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Params:", r.URL.Query())

	paramProject := r.URL.Query().Get("project")
	paramBranch := r.URL.Query().Get("branch")
	if paramProject == "" {
		fmt.Println("No proporciono el parametro PROJECT")
	}
	if paramBranch == "" {
		fmt.Println("No proporciono el parametro BRANCH")
	}

	xx := sonarqube.GetAllModules(w, r, paramProject)
	var ff sonarqube.MetricsComponentList
	var bb sonarqube.MetricsComponentList

	for _, element := range xx.Front.Projects {
		ff.Components = append(ff.Components, sonarqube.GetMetrics(w, r, element.Key, paramBranch))
	}

	for _, element := range xx.Back.Projects {
		bb.Components = append(bb.Components, sonarqube.GetMetrics(w, r, element.Key, paramBranch))
	}

	dat, err := ioutil.ReadFile("html/template.html")
	check(err)

	dataHtml := html.CreateHome(string(dat), xx)
	dataHtml = html.CreateBackend(dataHtml, bb, xx)
	dataHtml = html.CreateFrontend(dataHtml, ff, xx)

	var dir = ""
	if (paramProject == "bel")	{
		dir = "bel-personal/metrics/"
	}
	if (paramProject == "belcom")	{
		dir = "bel-comercial/metrics/"
	}

	t := time.Now()
	concatenated := fmt.Sprint("/apachebel/", dir, t.Year(), "-", int(t.Month()), "/report.html")

	newpath := filepath.Dir(concatenated)
	err = os.MkdirAll(newpath, os.ModePerm)

	er2 := ioutil.WriteFile(concatenated, []byte(dataHtml), 0644)

	check(er2)

	fmt.Fprint(w, "READY")
}

func metricsDaily(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Params:", r.URL.Query())

	paramProject := r.URL.Query().Get("project")
	paramBranch := r.URL.Query().Get("branch")
	if paramProject == "" {
		fmt.Println("No proporciono el parametro PROJECT")
	}
	if paramBranch == "" {
		fmt.Println("No proporciono el parametro BRANCH")
	}

	xx := sonarqube.GetAllModules(w, r, paramProject)
	var ff sonarqube.MetricsComponentList
	var bb sonarqube.MetricsComponentList

	for _, element := range xx.Front.Projects {
		ff.Components = append(ff.Components, sonarqube.GetMetrics(w, r, element.Key, paramBranch))
	}

	for _, element := range xx.Back.Projects {
		bb.Components = append(bb.Components, sonarqube.GetMetrics(w, r, element.Key, paramBranch))
	}

	dat, err := ioutil.ReadFile("html/template.html")
	check(err)

	dataHtml := html.CreateHome(string(dat), xx)
	dataHtml = html.CreateBackend(dataHtml, bb, xx)
	dataHtml = html.CreateFrontend(dataHtml, ff, xx)

	var dir = ""
	if (paramProject == "bel")	{
		dir = "bel-personal/metrics/"
	}
	if (paramProject == "belcom")	{
		dir = "bel-comercial/metrics/"
	}

	t := time.Now()
	concatenated := fmt.Sprint("/apachebel/", dir, t.Year(), "-", int(t.Month()), "/report.html")
	er2 := ioutil.WriteFile(concatenated, []byte(dataHtml), 0644)

	check(er2)

	fmt.Fprint(w, "READY")
}

//laisa04@cwpanama.net

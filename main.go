package main

import (
	"cualquier_vaina/sonarqube"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func getAllModules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Params:", r.URL.Query())

	param1 := r.URL.Query().Get("project")
  if param1 != "" {
    fmt.Println("No proporciono el parametro PROJECT")
  }

  xx := sonarqube.GetAllModules(w, r, param1)
	var ff sonarqube.MetricsComponentList
	var bb sonarqube.MetricsComponentList
	
	for _, element := range xx.Front.Projects {
		ff.Components = append(ff.Components, sonarqube.GetMetrics(w, r, element.Key))
	}

	for _, element := range xx.Back.Projects {
		bb.Components = append(bb.Components, sonarqube.GetMetrics(w, r, element.Key))
	}

	dat, err := ioutil.ReadFile("html/template.html")
	check(err)

	data := strings.ReplaceAll(string(dat), "{{mod_front}}", fmt.Sprint(len(xx.Front.Projects)))
	data = strings.ReplaceAll(data, "{{mod_back}}", fmt.Sprint(len(xx.Back.Projects)))
	total := (len(xx.Back.Projects) + len(xx.Front.Projects))
	data = strings.ReplaceAll(data, "{{mod_total}}", fmt.Sprint(total))

	var sonar_measures string
	var sonar_complexities string
	var filesT int
	var coverageT float64
	var codeSmellsT int
	var bugsT int
	var vulnerabilitiesT int
	var testT int
	var functionsT int
	var cognitiveComplexityT int
	var complexityT int
	for index, element := range bb.Components {
		var files string
		var coverage string
		var codeSmells string
		var bugs string
		var vulnerabilities string
		var test string
		var functions string
		var cognitiveComplexity string
		var complexity string
		for i := range element.Component.Measures {
			if element.Component.Measures[i].Metric == "bugs" {
				bugs = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(bugs)
				bugsT += i
			}
			if element.Component.Measures[i].Metric == "files" {
				files = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(files)
				filesT += i
			}
			if element.Component.Measures[i].Metric == "coverage" {
				coverage = element.Component.Measures[i].Value
				i, _ := strconv.ParseFloat(coverage, 64)
				coverageT += i
			}
			if element.Component.Measures[i].Metric == "code_smells" {
				codeSmells = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(codeSmells)
				codeSmellsT += i
			}
			if element.Component.Measures[i].Metric == "vulnerabilities" {
				vulnerabilities = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(vulnerabilities)
				vulnerabilitiesT += i
			}
			if element.Component.Measures[i].Metric == "tests" {
				test = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(test)
				testT += i
			}
			if element.Component.Measures[i].Metric == "functions" {
				functions = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(test)
				functionsT += i
			}
			if element.Component.Measures[i].Metric == "cognitive_complexity" {
				cognitiveComplexity = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(test)
				cognitiveComplexityT += i
			}
			if element.Component.Measures[i].Metric == "complexity" {
				complexity = element.Component.Measures[i].Value
				i, _ := strconv.Atoi(test)
				complexityT += i
			}
		}
		sonar_measures += fmt.Sprintf("<tr> <td class='text-center'>%v</td> <td class='text-left'>%v</td> <td class='text-center'>%v</td><td class='text-center'>%v</td><td class='text-right'>%v</td><td class='text-right'>%v</td>		<td class='text-right'>%v</td>		<td class='text-right'>%v</td></tr>", index+1, element.Component.Name, files, test, coverage, codeSmells, bugs, vulnerabilities)

		sonar_complexities += fmt.Sprintf("<tr> <td class='text-center'>%v</td> <td class='text-left'>%v</td> <td class='text-center'>%v</td><td class='text-right'>%v</td><td class='text-right'>%v</td>		<td class='text-right'>%v</td>		<td class='text-right'>%v</td></tr>", index+1, element.Component.Name, functions, complexity, test, cognitiveComplexity, bugs)
	}

	sonar_measuresT := fmt.Sprintf("<th class='text-right'>%v</th>	<th class='text-right'>%v</th>	<th class='text-right'>%.2f</th>	<th class='text-right'>%v</th>	<th class='text-right'>%v</th>	<th class='text-right'>%v</th>", filesT, testT, coverageT/float64(total), codeSmellsT, bugsT, vulnerabilitiesT)

	sonar_complexitiesT := fmt.Sprintf("<th class='text-right'>%v</th>	<th class='text-right'>%v</th>	<th class='text-right'>%.2f</th>	<th class='text-right'>%v</th>	<th class='text-right'>%v</th>", functionsT, complexityT, coverageT/float64(total), cognitiveComplexityT, bugsT)

	data = strings.ReplaceAll(data, "{{sonar_measures_rows}}", sonar_measures)
	data = strings.ReplaceAll(data, "{{sonar_measures}}", sonar_measuresT)
	data = strings.ReplaceAll(data, "{{sonar_complexities_rows}}", sonar_complexities)
	data = strings.ReplaceAll(data, "{{sonar_complexities}}", sonar_complexitiesT)

	er2 := ioutil.WriteFile("test.html", []byte(data), 0644)
	check(er2)
	// f, err := os.Create("test.html")
	// _, err2 := f.WriteString(tmpl)
	fmt.Fprint(w, "READY")
}

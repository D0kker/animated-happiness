package html

import (
	"cualquier_vaina/sonarqube"
	"fmt"
	"strings"
	"strconv"
	//"math"
)

type htmlDataStruct struct {
	sonar_measures       	string
	sonar_measuresT      	string
	sonar_complexities 		string
	sonar_complexitiesT   string
}

func CreateHome(htmlData string, xx sonarqube.ProjectListModules) string {
  htmlData = strings.ReplaceAll(string(htmlData), "{{mod_front}}", fmt.Sprint(len(xx.Front.Projects)))
	htmlData = strings.ReplaceAll(htmlData, "{{mod_back}}", fmt.Sprint(len(xx.Back.Projects)))
	total := (len(xx.Back.Projects) + len(xx.Front.Projects))
	htmlData = strings.ReplaceAll(htmlData, "{{mod_total}}", fmt.Sprint(total))

	return htmlData
}

func getData(xx sonarqube.MetricsComponentList, pp sonarqube.ProjectListModules) htmlDataStruct {
	total := (len(pp.Back.Projects) + len(pp.Front.Projects))

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

	for index, element := range xx.Components {
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

		sonar_measures += fmt.Sprintf("<tr> <td class='text-left'>%v</td> <td class='text-left'>%v</td> <td class='text-right'>%v</td><td class='text-right'>%v</td><td class='text-right'>%v</td><td class='text-right'>%v</td>		<td class='text-right'>%v</td>		<td class='text-right'>%v</td></tr>", index+1, element.Component.Name, files, test, coverage, codeSmells, bugs, vulnerabilities)

		parsedFunctions, _ := strconv.Atoi(functions)
		parsedcognitiveComplexity, _ := strconv.Atoi(cognitiveComplexity)
		parsedcomplexity, _ := strconv.Atoi(complexity)

		var complexityAVG float64
		if complexityAVG = 0; parsedFunctions > 0 {
			complexityAVG = float64(parsedcomplexity)/float64(parsedFunctions)
			//complexityAVG = math.Round(complexityAVG)
		}

		var cognitiveComplexityAVG float64
		if cognitiveComplexityAVG = 0; parsedFunctions > 0 {
		  cognitiveComplexityAVG = float64(parsedcognitiveComplexity)/float64(parsedFunctions)
			//cognitiveComplexityAVG = math.Round(cognitiveComplexityAVG)
		}

		sonar_complexities += fmt.Sprintf("<tr> <td class='text-left'>%v</td> <td class='text-left'>%v</td> <td class='text-right'>%v</td><td class='text-right'>%v</td><td class='text-right'>%f</td>		<td class='text-right'>%v</td>		<td class='text-right'>%f</td></tr>", index+1, element.Component.Name, functions, complexity, complexityAVG, cognitiveComplexity, cognitiveComplexityAVG)
	}

	sonar_measuresT := fmt.Sprintf("<th class='text-right'>%v</th>	<th class='text-right'>%v</th>	<th class='text-right'>%f</th>	<th class='text-right'>%v</th>	<th class='text-right'>%v</th>	<th class='text-right'>%v</th>", filesT, testT, coverageT/float64(total), codeSmellsT, bugsT, vulnerabilitiesT)

	//parsedFunctionsT, _ := functionsT
	// parsedcognitiveComplexityT, _ := strconv.Atoi(cognitiveComplexityT)
	// parsedcomplexityT, _ := strconv.Atoi(complexityT)

	var complexityTAVG float64
	if complexityTAVG = 0; functionsT > 0 {
		complexityTAVG = float64(complexityT)/float64(functionsT)
		//complexityTAVG = math.Round(complexityTAVG)
	}

	var cognitiveComplexityTAVG float64
	if cognitiveComplexityTAVG = 0; functionsT > 0 {
		cognitiveComplexityTAVG = float64(cognitiveComplexityT)/float64(functionsT)
		//cognitiveComplexityTAVG = math.Round(cognitiveComplexityTAVG)
	}

	sonar_complexitiesT := fmt.Sprintf("<th class='text-right'>%v</th>	<th class='text-right'>%v</th>	<th class='text-right'>%.2f</th>	<th class='text-right'>%v</th>	<th class='text-right'>%.2f</th>", functionsT, complexityT, complexityTAVG, cognitiveComplexityT, cognitiveComplexityTAVG)

	return htmlDataStruct{
		sonar_measures: sonar_measures,
		sonar_measuresT: sonar_measuresT,
		sonar_complexities: sonar_complexities,
		sonar_complexitiesT: sonar_complexitiesT,
	}
}

func CreateBackend(htmlData string, xx sonarqube.MetricsComponentList, pp sonarqube.ProjectListModules) string {
	htmlDataStruct := getData(xx, pp)

	htmlData = strings.ReplaceAll(htmlData, "{{sonar_measures_rows_BB}}", htmlDataStruct.sonar_measures)
	htmlData = strings.ReplaceAll(htmlData, "{{sonar_measures_BB}}", htmlDataStruct.sonar_measuresT)
	htmlData = strings.ReplaceAll(htmlData, "{{sonar_complexities_rows_BB}}", htmlDataStruct.sonar_complexities)
	htmlData = strings.ReplaceAll(htmlData, "{{sonar_complexities_BB}}", htmlDataStruct.sonar_complexitiesT)

	return htmlData
}

func CreateFrontend(htmlData string, xx sonarqube.MetricsComponentList, pp sonarqube.ProjectListModules) string {
	htmlDataStruct := getData(xx, pp)

	htmlData = strings.ReplaceAll(htmlData, "{{sonar_measures_rows_FF}}", htmlDataStruct.sonar_measures)
	htmlData = strings.ReplaceAll(htmlData, "{{sonar_measures_FF}}", htmlDataStruct.sonar_measuresT)
	htmlData = strings.ReplaceAll(htmlData, "{{sonar_complexities_rows_FF}}", htmlDataStruct.sonar_complexities)
	htmlData = strings.ReplaceAll(htmlData, "{{sonar_complexities_FF}}", htmlDataStruct.sonar_complexitiesT)

	return htmlData
}
package sonarqube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MetricsComponent struct {
	Component struct {
		ID          string `json:"id"`
		Key         string `json:"key"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Qualifier   string `json:"qualifier"`
		Measures    []struct {
			Metric    string `json:"metric"`
			Value     string `json:"value"`
			BestValue bool   `json:"bestValue,omitempty"`
		} `json:"measures"`
	} `json:"component"`
	Metrics []struct {
		Key                   string `json:"key"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		Domain                string `json:"domain"`
		Type                  string `json:"type"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter"`
		Qualitative           bool   `json:"qualitative"`
		Hidden                bool   `json:"hidden"`
		Custom                bool   `json:"custom"`
		BestValue             string `json:"bestValue,omitempty"`
	} `json:"metrics"`
	Period struct {
		Index     int    `json:"index"`
		Mode      string `json:"mode"`
		Parameter string `json:"parameter"`
	} `json:"period"`
}

type MetricsComponentList struct {
	Components []MetricsComponent
}

func GetMetrics(w http.ResponseWriter, r *http.Request, projectKey string, paramBranch string) MetricsComponent {
	//fmt.Println(projectKey)
	var bird MetricsComponent
	metrics := "ncloc,complexity,violations,files,coverage,code_smells,bugs,vulnerabilities,cognitive_complexity,functions,tests"
	route := fmt.Sprintf("http://devops/sonar/api/measures/component?additionalFields=period,metrics&branch=%s&component=%s&metricKeys=%s", paramBranch, projectKey, metrics)

	fmt.Println(route)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", route, nil)
	req.Header.Set("Authorization", "Basic ZmQxMTU5MGFmZDIyMDQ3Y2YzNzhjYjczYWMzYmJlYzAzOTNjMWZjMTo=")

	res, _ := client.Do(req)

	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if !strings.Contains(string(responseData), "errors"){
		json.Unmarshal([]byte(responseData), &bird)
	}
	//fmt.Fprint(w, string(responseData))
	//fmt.Println(string(responseData))

	return bird
}

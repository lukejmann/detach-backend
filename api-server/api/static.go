package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	m "github.com/lukejmann/detach2-backend/api-server/models"
)

//sessionID is "NA" if invalid
func serveFetchAppDomains(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("serving fetch app domains\n")
	read, _ := ioutil.ReadFile("./static/prod/appDomains.json")
	var appDomains []m.AppDomain
	err := json.Unmarshal(read, &appDomains)
	// pp.Println("appDomains: ", appDomains)
	if err != nil {
		fmt.Printf("error decoding app domains: %v\n", err.Error())
	}

	if len(appDomains) > 0 {
		fmt.Printf("API successfully fetched appDomains")
		// w.WriteHeader(http.Stasuc)
	} else {
		fmt.Printf("API failed to fetch appDomains. length is zero")
		// w.WriteHeader(http.StatusNotFound)
	}
	return writeJSON(w, appDomains)

}

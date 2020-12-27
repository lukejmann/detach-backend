package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/k0kubun/pp"

	m "github.com/lukejmann/detach2-backend/api-server/models"
)

//sessionID is "NA" if invalid
func serveCreateSession(w http.ResponseWriter, r *http.Request) error {

	// var s string
	// err := json.NewDecoder(r.Body).Decode(&s)
	// if err != nil {
	// 	return err
	// }

	var opt m.SessionCreateOpt
	err := json.NewDecoder(r.Body).Decode(&opt)
	if err != nil {
		return err
	}

	store, err := GetDatastore()
	if err != nil {
		return err
	}

	res, err := store.Sessions.Create(opt)
	pp.Println("res:", res)
	if res.Success {
		fmt.Printf("API created session. sessionID: %v\n", res.SessionID)
		w.WriteHeader(http.StatusCreated)
	} else {
		fmt.Printf("API failed to create session. sessionID: %v. Error: %v\n", res.SessionID, err)
		w.WriteHeader(http.StatusNotModified)
	}
	e := writeJSON(w, res)
	fmt.Printf("Error in writeJSON: %v\n", e)
	return e

}

//sessionID is "NA" if invalid
func serveCancelSession(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("in serveCancelSession")
	var opt m.SessionCancelOpt
	err := json.NewDecoder(r.Body).Decode(&opt)
	if err != nil {
		return err
	}

	store, err := GetDatastore()
	if err != nil {
		return err
	}

	success, err := store.Sessions.Cancel(opt)
	if success {
		fmt.Printf("API canceled session. sessionID: %v\n", opt.SessionID)
		// w.WriteHeader(http.StatusOK)
	} else {
		fmt.Printf("API failed to cancel session. sessionID: %v\n", opt.SessionID)
		// w.WriteHeader(http.StatusNotModified)
	}
	fmt.Printf("serveCancelSession success: %v\n", success)
	return writeJSON(w, success)
}

package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lukejmann/detach2-backend/api-server/router"

	"github.com/gorilla/mux"
)

func setupAPI() {
	GetDatastore()
}

func Handler() *mux.Router {
	setupAPI()
	m := router.API()

	m.Get(router.Login).Handler(handler(serveLogin))

	m.Get(router.CreateSession).Handler(handler(serveCreateSession))
	m.Get(router.CancelSession).Handler(handler(serveCancelSession))

	// m.Get(router.CheckReceipt).Handler(handler(serveCheckReceipt))

	m.Get(router.FetchAppDomains).Handler(handler(serveFetchAppDomains))

	return m
}

type handler func(http.ResponseWriter, *http.Request) error

//I believe this just handles errors that come from calling the handler
// (it just prints the error to the HTML page)
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in ServeHTTP")
	err := h(w, r)
	// pp.Println("err:", err)
	if err != nil {
		fmt.Printf("err with request %v\nerr:%v\n", r.URL.Path, err)
		fmt.Fprintf(w, "error: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

package router

import (
	"github.com/gorilla/mux"
)

func API() *mux.Router {
	m := mux.NewRouter()
	// STATIC_DIR := "/detach/static/prod"

	m.Path("/users/login").Methods("GET").Queries("userID", "{userID}", "email", "{email}").Name(Login)
	// m.Path("/users/checkReceipt").Methods("POST").Name(CheckReceipt)

	m.Path("/sessions/create").Methods("POST").Name(CreateSession)
	m.Path("/sessions/cancel").Methods("POST").Name(CancelSession)

	m.Path("/static/ADs").Methods("GET").Name(FetchAppDomains)

	// m.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
	// 	http.FileServer(http.Dir(STATIC_DIR))))

	return m
}

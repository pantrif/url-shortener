package main

import (
    "log"
    "net/http"
	"database/sql"
    "github.com/gorilla/mux"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Init() {
	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/{hash}", redirect).Methods("GET")
    a.Router.HandleFunc("/shorten", shorten).Methods("POST")
 }

 func (a *App) Run(port string) {
    log.Fatal(http.ListenAndServe(port, a.Router))
 }

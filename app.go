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
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/{hash}", a.Redirect).Methods("GET")
    a.Router.HandleFunc("/shorten", a.Shorten).Methods("POST")
 }

 func (a *App) Run(port string) {
    log.Fatal(http.ListenAndServe(port, a.Router))
 }

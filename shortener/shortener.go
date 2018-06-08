package shortener

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//App with a router and db as dependencies
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Init routes
func (a *App) Init() {
	a.Router.HandleFunc("/", a.Home).Methods("GET")
	a.Router.HandleFunc("/{hash}", a.Redirect).Methods("GET")
	a.Router.HandleFunc("/shorten", a.Shorten).Methods("POST")
}

//Run the app
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

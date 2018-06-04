package main

import (
	"os"
	"github.com/gorilla/mux"
)

func main() {
	a := App{}
	a.DB = InitDB("mysql")
	a.Router = mux.NewRouter()
	a.Init()
	a.Run(":"+os.Getenv("EXPOSED_PORT"))
}


package main

import (
	"os"

	"github.com/pantrif/url-shortener/shortener"

	"github.com/gorilla/mux"
)

func main() {
	a := shortener.App{}
	a.DB = shortener.InitDB("mysql")
	a.Router = mux.NewRouter()
	a.Init()
	a.Run(":" + os.Getenv("EXPOSED_PORT"))
}

package main

import (
	"os"

	"github.com/pantrif/url-shortener/shortener"

	"github.com/gorilla/mux"
)

func serve(a shortener.App) {
	a.Router = mux.NewRouter()
	a.Init()
	a.Run(":" + os.Getenv("EXPOSED_PORT"))
}

func main() {
	a := shortener.App{}
	a.DB = shortener.InitDB("mysql")
	serve(a)
}

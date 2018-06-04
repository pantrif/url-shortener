package main

import (
	"net/http"
	"net/url"
	"github.com/gorilla/mux"
	"github.com/bitly/go-simplejson"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here"))
}

func shorten(w http.ResponseWriter, r *http.Request) {
	var id int64
	url := r.FormValue("url")

	json := simplejson.New()

	if (!isValidUrl(url)){
		w.WriteHeader(http.StatusBadRequest)
		json.Set("error", "Invalid url")
		sendResponse(json, w)
		return;
	}

	db.QueryRow("SELECT id FROM shortened_urls WHERE long_url = ?", url).Scan(&id)
	if (id == 0){
		res, err := db.Exec(`INSERT INTO shortened_urls (long_url, created) VALUES(?, now())`, url)
		if err != nil {
			log.Println("Error:", err.Error())
		}
		id, _ = res.LastInsertId()
	}
	hash := encode(id)
	json.Set("url", r.Host  + "/" + hash)
	sendResponse(json, w)
}

func redirect(w http.ResponseWriter, r *http.Request)  {
	var id int64
	var long_url string

	json := simplejson.New()

	vars := mux.Vars(r)
	hash := vars["hash"]

	decoded_id := decode(hash)
	db.QueryRow("SELECT id, long_url FROM shortened_urls WHERE id = ?", decoded_id).Scan(&id, &long_url)

	if (id == 0){
		w.WriteHeader(http.StatusNotFound)
		json.Set("error", "Not found")
		sendResponse(json, w)
		return;
	}
	http.Redirect(w, r, long_url, http.StatusSeeOther)
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

func sendResponse(json *simplejson.Json, w http.ResponseWriter)  {
	w.Header().Set("Content-Type", "application/json")
	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Write(payload)
}
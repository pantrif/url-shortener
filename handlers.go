package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type Body struct {
	Url string `json:"url"`
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here"))
}

func (a *App) Shorten(w http.ResponseWriter, r *http.Request) {
	var id int64
	var body Body

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	url := body.Url

	if !isValidUrl(url) {
		respondWithError(w, http.StatusBadRequest, "Invalid url")
		return
	}

	a.DB.QueryRow("SELECT id FROM shortened_urls WHERE long_url = ?", url).Scan(&id)
	if id == 0 {
		res, err := a.DB.Exec(`INSERT INTO shortened_urls (long_url, created) VALUES(?, now())`, url)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		id, _ = res.LastInsertId()
	}
	hash := encode(id)
	body.Url = r.Host + "/" + hash
	sendResponse(w, http.StatusOK, body)
}

func (a *App) Redirect(w http.ResponseWriter, r *http.Request) {
	var id int64
	var long_url string

	vars := mux.Vars(r)
	hash := vars["hash"]

	decoded_id := decode(hash)
	a.DB.QueryRow("SELECT id, long_url FROM shortened_urls WHERE id = ?", decoded_id).
		Scan(&id, &long_url)

	if id == 0 {
		respondWithError(w, http.StatusNotFound, "Not found")
		return
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

func respondWithError(w http.ResponseWriter, code int, message string) {
	sendResponse(w, code, map[string]string{"error": message})
}

func sendResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

package shortener

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

//Body is the response body
type Body struct {
	URL string `json:"url"`
}

//Home page
func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here"))
}

//Shorten url POST method
func (a *App) Shorten(w http.ResponseWriter, r *http.Request) {
	var id int64
	var body Body

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	url := body.URL

	if !isValidURL(url) {
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
	body.URL = r.Host + "/" + hash
	sendResponse(w, http.StatusOK, body)
}

//Redirect route
func (a *App) Redirect(w http.ResponseWriter, r *http.Request) {
	var id int64
	var longURL string

	vars := mux.Vars(r)
	hash := vars["hash"]

	decodedID := decode(hash)
	a.DB.QueryRow("SELECT id, long_url FROM shortened_urls WHERE id = ?", decodedID).
		Scan(&id, &longURL)

	if id == 0 {
		respondWithError(w, http.StatusNotFound, "Not found")
		return
	}
	http.Redirect(w, r, longURL, http.StatusSeeOther)
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true

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

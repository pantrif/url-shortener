package shortener

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var a App

func TestHomeRedirect(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	site := "http://www.google.com"

	a = App{}
	a.Router = mux.NewRouter()
	a.DB = db
	a.Init()
	req, _ := http.NewRequest("GET", "/", nil)
	r := executeRequest(req)

	assertEqual(t, r.Body.String(), "Nothing to see here")
	assertEqual(t, r.Code, 200)

	req, _ = http.NewRequest("POST", "/", nil)
	r = executeRequest(req)
	assertEqual(t, r.Code, 405)

	rows := sqlmock.NewRows([]string{"id", "long_url"}).
		AddRow(1722946, site)

	mock.ExpectQuery("SELECT id, long_url FROM shortened_urls").WithArgs(1722946).WillReturnRows(rows)
	req, _ = http.NewRequest("GET", "/hash", nil)
	r = executeRequest(req)
	assertEqual(t, r.Code, 303)

	mock.ExpectQuery("SELECT id, long_url FROM shortened_urls").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "long_url"}))
	req, _ = http.NewRequest("GET", "/hash", nil)
	r = executeRequest(req)
	assertEqual(t, r.Code, 404)
}

func TestShorten(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	site := "http://www.google.com"
	invalidSite := "trash"

	a = App{}
	a.Router = mux.NewRouter()
	a.DB = db
	a.Init()

	req, _ := http.NewRequest("POST", "/shorten", strings.NewReader(""))
	r := executeRequest(req)
	assertEqual(t, r.Code, 400)

	jsonReq := "{\"url\":\"" + invalidSite + "\"}"
	req, _ = http.NewRequest("POST", "/shorten", strings.NewReader(jsonReq))
	r = executeRequest(req)
	assertEqual(t, r.Code, 400)

	jsonReq = "{\"url\":\"" + site + "\"}"
	mock.ExpectQuery("SELECT id FROM shortened_urls").WithArgs(site).WillReturnRows(sqlmock.NewRows([]string{"id"}))
	mock.ExpectExec("INSERT INTO shortened_urls").WithArgs(site).WillReturnError(fmt.Errorf("an error occurred"))
	req, _ = http.NewRequest("POST", "/shorten", strings.NewReader(jsonReq))
	r = executeRequest(req)
	assertEqual(t, r.Code, 500)

	jsonReq = "{\"url\":\"" + site + "\"}"

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1722946)
	mock.ExpectQuery("SELECT id FROM shortened_urls").WithArgs(site).WillReturnRows(rows)
	mock.ExpectExec("INSERT INTO shortened_urls").WithArgs(site).WillReturnResult(sqlmock.NewResult(1, 1))
	req, _ = http.NewRequest("POST", "/shorten", strings.NewReader(jsonReq))
	r = executeRequest(req)
	assertEqual(t, r.Body.String(), "{\"url\":\"/g_sh\"}")
	assertEqual(t, r.Code, 200)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

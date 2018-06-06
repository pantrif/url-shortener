package main_test

import(
	"net/http/httptest"
	"testing"
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"."
)

var a main.App

func TestMain(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	a = main.App{}
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
		AddRow(1722946, "http://www.google.com")

	mock.ExpectQuery("SELECT id, long_url FROM shortened_urls").WithArgs(1722946).WillReturnRows(rows)
	req, _ = http.NewRequest("GET", "/hash", nil)
	r = executeRequest(req)
	assertEqual(t, r.Code, 303)

	mock.ExpectQuery("SELECT id, long_url FROM shortened_urls").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "long_url"}))
	req, _ = http.NewRequest("GET", "/hash", nil)
	r = executeRequest(req)
	assertEqual(t, r.Code, 404)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    a.Router.ServeHTTP(rr, req)

    return rr
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
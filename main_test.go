package main

import (
	"net/http"

	"testing"
	"time"

	"github.com/pantrif/url-shortener/shortener"
)

var a shortener.App

func TestMain(t *testing.T) {
	a := shortener.App{}
	go serve(a)

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://127.0.0.1:8081/", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assertEqual(t, http.StatusOK, resp.StatusCode)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

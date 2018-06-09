package main

import (
	"net/http"

	"testing"
	"time"

	"github.com/json-iterator/go/assert"
	"github.com/pantrif/url-shortener/shortener"
)

var a shortener.App

func TestMain(t *testing.T) {
	a := shortener.App{}
	go serve(a)

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	r, _ := http.NewRequest("GET", "http://localhost:8081/", nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

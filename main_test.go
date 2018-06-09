package main

import (
	"testing"

	"github.com/pantrif/url-shortener/shortener"
)

var a shortener.App

func TestMain(t *testing.T) {
	a := shortener.App{}
	go serve(a)
}

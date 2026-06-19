package main

import (
	"fmt"
	"net/http"
)

func v2handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to version 2")
}

func legacyHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

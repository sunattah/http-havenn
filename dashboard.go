package main

import (
	"fmt"
	"net/http"
)

func dashboardhandler(w http.ResponseWriter, r *http.Request) {

	apikey := r.Header.Get("X-API-Key")
	if apikey != "secret123" {
		http.Error(w, "http.StatusUnauthorized", http.StatusUnauthorized)
		return

	}
	fmt.Fprintf(w, "Welcome")

}

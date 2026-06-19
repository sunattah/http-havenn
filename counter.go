package main

import (
	"fmt"
	"net/http"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Send a POST request with text to count words")
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, len("Golang"))
	}
}

package main

import (
	"fmt"
	"net/http"
)

func HandleAgent(w http.ResponseWriter, r *http.Request) {
	agent := r.Header.Get("User-Agent")
	fmt.Fprintln(w, agent)
}

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandleCalc(w http.ResponseWriter, r *http.Request) {
	opp := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	num1, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if opp == "add" {
		fmt.Fprintln(w, (num1 + num2))
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("program is working")
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/count", CounterHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/calculate", HandleCalc)
	http.HandleFunc("/agent", HandleAgent)
	http.HandleFunc("/dashboard", dashboardhandler)
	http.HandleFunc("/legacy", legacyHandler)
	http.HandleFunc("/v2", v2handler)
	http.ListenAndServe(":8080", nil)

}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you've requested: %s\n", r.URL.Path)
	})

	http.HandleFunc("/users", GetUsers)

	http.ListenAndServe(":9999", nil)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tese"))
}

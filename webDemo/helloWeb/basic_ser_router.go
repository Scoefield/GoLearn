package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

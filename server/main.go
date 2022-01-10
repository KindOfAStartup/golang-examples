package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hik")
	})

	http.HandleFunc("/test", handleloadJson)
	http.HandleFunc("/spotify/albums", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "responses/albums.json")
	})
	http.HandleFunc("/nba/teams", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "responses/nbateams.json")
	})
	log.Fatal(http.ListenAndServe(":8099", nil))

}

func handleloadJson(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "responses/test.json")
}

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
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func handleloadJson(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "responses/test.json")
}

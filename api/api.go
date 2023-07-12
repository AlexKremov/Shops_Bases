package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/search", handleSearch)

	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	results := performSearch(query, lat, long)

	fmt.Fprintf(w, "Результаты поиска: %v", results)
}

func performSearch(query, lat, long string) []string {

	searchResult := fmt.Sprintln(query, lat, long)
	return []string{searchResult}
}

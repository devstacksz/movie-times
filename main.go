package main

import (
	"fmt"
)

func main() {
	movies := getPopularMovies()
	fmt.Println("Movies:", movies)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, render.RenderTemplate(movies))
	// })

	// http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
	// 	dat, err := os.ReadFile("styles.css")
	// 	check(err)
	// 	w.Header().Set("Content-Type", "text/css")
	// 	fmt.Fprint(w, string(dat))
	// })

	// fmt.Println("Server running on http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
}

// TODO: Make the scrape package be importable

// TODO: Make a commit and does not use channels
// TODO: Figure out how to listen to the channel for movies

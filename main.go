package main

import (
	"fmt"
	"movie-times/render"
	"movie-times/scrape"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func host() {
	// Register route handlers
	http.HandleFunc("/", indexHandler)

	// Start the HTTP server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	// Grab all the data first
	movies := scrape.Scrape()
	fmt.Println("Movies:", movies)

	// Render the template using the data
	// Need to pass in the data from scraping

	// Start the server
	// Register route handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output := render.RenderTemplate(movies)
		fmt.Fprint(w, output)
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

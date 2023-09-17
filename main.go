package main

import (
	"fmt"
	"movie-times/render"
	"movie-times/scrape"
	"net/http"
)

func main() {
	movies := scrape.Scrape()
	fmt.Println("Movies:", movies)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, render.RenderTemplate(movies))
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"movie-times/render"
	"movie-times/scrape"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	movies := scrape.Scrape()
	fmt.Println("Movies:", movies)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, render.RenderTemplate(movies))
	})

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		dat, err := os.ReadFile("styles.css")
		check(err)
		w.Header().Set("Content-Type", "text/css")
		fmt.Fprint(w, string(dat))
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

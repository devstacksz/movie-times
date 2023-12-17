package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type Movie struct {
	Title       string
	Date        time.Time
	Description string
	PosterPath  string
}

func formatYear(t time.Time) string {
	return t.Format("January 2, 2006")
}
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	limit := 25
	movieChan := make(chan Movie, limit)
	movies := make([]Movie, limit)
	c := colly.NewCollector()

	go getPopularMovies(c, movieChan)
	c.Wait()

	for count := 0; count < limit; count++ {
		movies[count] = <-movieChan
	}

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		t := template.New("movie-table.html").Funcs(template.FuncMap{
			"formatYear": formatYear,
		})

		tmpl, err := t.ParseFiles("./templates/movie-table.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

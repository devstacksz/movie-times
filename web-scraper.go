package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/gocolly/colly"
)

const rtUrl = "https://www.rottentomatoes.com/browse/movies_in_theaters/sort:popular"

func getLandmarkUrl(location string) string {
	urlParts := []string{"https://www.landmarkcinemas.com/showtimes/", location, "/"}
	return strings.Join(urlParts, "")
}

func getLandmarkShowtimes(c *colly.Collector) {
	// TODO: Implement
}

func getPopularMovies(c *colly.Collector, movieChan chan Movie) {
	tmdbClient, err := tmdb.Init(os.Getenv("TMDB_KEY"))
	if err != nil {
		fmt.Println(err)
	}

	c.OnHTML("a[data-track]", func(e *colly.HTMLElement) {
		e.DOM.Find("span.p--small").Each(func(index int, item *goquery.Selection) {

			title := strings.TrimSpace(item.Text())

			e.DOM.Find("span.p--small + span").Each(func(index int, item *goquery.Selection) {
				stringDate := strings.TrimPrefix(strings.TrimSpace(item.Text()), "Opened ")
				stringDate = strings.TrimPrefix(stringDate, "Opens ")

				date, err := time.Parse("Jan 2, 2006", stringDate)
				if err != nil {
					log.Fatal(err)
				}

				go createMovie(tmdbClient, title, date, movieChan)
			})
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	err = c.Visit(rtUrl)

	if err != nil {
		log.Fatal(err)
	}
}

func createMovie(tmdbClient *tmdb.Client, title string, date time.Time, movieChan chan Movie) {
	search, err := tmdbClient.GetSearchMovies(title, nil)
	if err != nil {
		fmt.Println(err)
	}

	if len(search.Results) == 0 {
		fmt.Println("No movie found")
		return
	}

	movie := Movie{
		Title:       title,
		Date:        date,
		Description: search.Results[0].Overview,
		PosterPath:  search.Results[0].PosterPath,
	}

	movieChan <- movie

}

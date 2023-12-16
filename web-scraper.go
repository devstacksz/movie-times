package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

const rtUrl = "https://www.rottentomatoes.com/browse/movies_in_theaters/sort:popular"

func getPopularMovies() []string {
	c := colly.NewCollector()
	var movies []string

	c.OnHTML("a[data-track]", func(e *colly.HTMLElement) {
		e.DOM.Find("span.p--small").Each(func(index int, item *goquery.Selection) {
			movies = append(movies, strings.TrimSpace(item.Text())+"\n")
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	err := c.Visit(rtUrl)
	if err != nil {
		log.Fatal(err)
	}

	c.Wait()
	return movies
}

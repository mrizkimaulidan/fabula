package main

import (
	"flag"
	"strings"
)

var username string

func main() {
	flag.StringVar(&username, "username", "", "-username the instagram username")
	flag.Parse()

	i := Instagram{Username: username}
	parser := NewParser(&i)

	response := parser.Call()

	scraper := NewScraper(&i)
	scraper.Scrape(strings.NewReader(response.HTML.(string)))
}

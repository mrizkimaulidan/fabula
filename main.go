package main

import (
	"flag"
	"strings"
)

var username string

func main() {
	flag.StringVar(&username, "username", "", "-username=john.doe")
	flag.Parse()

	p := NewParser()
	response := p.Parse(username)

	s := NewScraper()
	s.Request.File.SetInstagramUsername(username)
	s.Scrape(strings.NewReader(response.HTML))
}

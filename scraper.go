package main

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	File      *File
	Instagram *Instagram
	Log       *log.Logger
}

// Core of this application. Do scraping to igpanda and
// download the story. doc.Find() will looping until
// how many the storyURL has found.
func (sc *Scraper) Scrape(r io.Reader) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		sc.Log.Fatalf("error reading from reader %s", err)
	}

	sc.Log.Println("starting to scrape story from", sc.Instagram.Username)
	doc.Find(".post-wrapper .download-button").Each(func(i int, s *goquery.Selection) {
		storyURL, _ := s.Attr("href") // getting the story URL

		sc.File.ShowDownloadText()
		sc.File.Download(storyURL)
		sc.File.IncrementDownloadCount()
	})

	sc.Log.Printf("%d story downloaded from %s", sc.File.GetDownloadCount(), sc.Instagram.Username)
}

func NewScraper(i *Instagram) *Scraper {
	return &Scraper{
		File:      NewFile(i),
		Instagram: i,
		Log:       log.Default(),
	}
}

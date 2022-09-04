package main

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	Request *Request
	Logger  *log.Logger
}

func (sc *Scraper) Scrape(r io.Reader) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		sc.Logger.Fatalln("error reading the reader", err)
	}

	sc.Logger.Println("starting to scrape story from", sc.Request.File.InstagramUsername)
	doc.Find(".post-wrapper .download-button").Each(func(i int, s *goquery.Selection) {
		storyURL, _ := s.Attr("href")

		extension := sc.Request.File.GetFileExtension(storyURL)
		sc.Request.File.SetExtension(extension)

		sc.Request.ShowDownloadText()
		sc.Request.Download(storyURL)
		sc.Request.IncrementDownloadCount()
	})

	sc.Request.DownloadedSuccessText()
}

func NewScraper() *Scraper {
	return &Scraper{
		Request: NewRequest(),
		Logger:  log.Default(),
	}
}

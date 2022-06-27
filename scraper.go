package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

// the scrapped website URL
var webURL = "https://hookgram.com/en/u/%s/stories"

type Scraper struct {
	*colly.Collector
	*Request
	FileExtension          string
	InstagramUsername      string
	DownloadedStoryCounter int    // for knowing how many story are has been downloaded
	FileURL                string // file url is a the url of the instagram story content that found when scraping the web
	LoadingText            string
	*log.Logger
}

func NewScraper(instagramUsername string) *Scraper {
	return &Scraper{
		Collector:         colly.NewCollector(),
		Request:           NewRequest(),
		InstagramUsername: instagramUsername,
		Logger:            log.Default(),
		LoadingText:       "downloading, please wait",
	}
}

// Checking the story content extension
// if the story content are videos, change to mp4
// if the story content are images, change to jpg
func (s *Scraper) checkExtension(str string) {
	if strings.Contains(str, "VIDEO") {
		s.FileExtension = "mp4"
	} else {
		s.FileExtension = "jpg"
	}
}

// Showing the loading text
func (s *Scraper) showLoadingText() {
	s.LoadingText += "."
	s.Logger.Println(s.LoadingText)
}

// Increment the downloaded story counter
func (s *Scraper) incrementDownloadedStoryCounter() {
	s.DownloadedStoryCounter++
}

// This function run scraping to the target website url
// basically this is the core of this application
func (s *Scraper) Scraping() {
	s.Collector.OnHTML("a.card.lift.h-100", func(h *colly.HTMLElement) {
		s.FileURL = h.Attr("href")

		s.incrementDownloadedStoryCounter()
		s.checkExtension(h.Text)

		s.showLoadingText()

		// downloading instagram story content
		s.Request.DoRequest(s.InstagramUsername, s.FileURL, s.FileExtension)
	})

	s.Collector.OnRequest(func(r *colly.Request) {
		s.Logger.Printf("starting to scrape story from %s", s.InstagramUsername)
	})

	s.Collector.OnScraped(func(r *colly.Response) {
		s.Logger.Printf("%d story downloaded from %s", s.DownloadedStoryCounter, s.InstagramUsername)
	})

	// visit the website URL
	s.Collector.Visit(fmt.Sprintf(webURL, s.InstagramUsername))
}

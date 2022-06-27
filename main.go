package main

func main() {
	f := NewFlag()
	args := f.ParsingFlag()

	s := NewScraper(args["username"].(string))
	s.Scraping()
}

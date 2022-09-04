package main

import (
	"log"
	"net/http"
)

type Request struct {
	File            *File
	DownloadCount   int
	DownloadingText string
	Logger          *log.Logger
}

// Show the downloaded success text.
func (r *Request) DownloadedSuccessText() {
	r.Logger.Printf("%d story downloaded from %s", r.DownloadCount, r.File.InstagramUsername)
}

// Show the downloading text process.
func (r *Request) ShowDownloadText() {
	r.DownloadingText += "."
	r.Logger.Println(r.DownloadingText)
}

// Increment the download count for tracking
// how many has been downloaded.
func (r *Request) IncrementDownloadCount() {
	r.DownloadCount++
}

// This function downloading the story file.
// Downloading and then do saving to filesystem.
func (r *Request) Download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		r.Logger.Fatalln("error requesting", err)
	}
	defer resp.Body.Close()

	err = r.File.CreateDir()
	if err != nil {
		r.Logger.Fatalln("error creating directory", err)
	}

	file, err := r.File.CreateFile()
	if err != nil {
		r.Logger.Fatalln("error creating file", err)
	}
	defer file.Close()

	err = r.File.CopyFile(file, resp.Body)
	if err != nil {
		r.Logger.Fatalln("error copying file", err)
	}
}

func NewRequest() *Request {
	return &Request{
		DownloadingText: "downloading, please wait",
		File:            NewFile(),
		Logger:          log.Default(),
	}
}

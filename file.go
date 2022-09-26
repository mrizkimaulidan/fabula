package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type File struct {
	Extension     string
	Instagram     *Instagram
	DownloadCount int // count how many story has been downloaded
	Log           *log.Logger
}

// Increment the download count
func (f *File) IncrementDownloadCount() {
	f.DownloadCount++
}

// Get the download count
func (f *File) GetDownloadCount() int {
	return f.DownloadCount
}

// Show downloading text status
func (f *File) ShowDownloadText() {
	f.Log.Println("downloading, please wait")
}

// Get the file extension from URL.
// The URL is the full URL of the story file, based on
// the scraped result
func (f *File) GetExtension(url string) string {
	fullURL := url
	s := strings.Split(fullURL, ".")
	extension := s[len(s)-1]

	return extension
}

// Download the story and save the story to the disk.
func (f *File) Download(URL string) {
	resp, err := http.Get(URL)
	if err != nil {
		f.Log.Fatalf("error requesting to URL %v", err)
	}
	defer resp.Body.Close()

	err = os.MkdirAll(fmt.Sprintf("stories/%s", f.Instagram.Username), os.ModePerm)
	if err != nil {
		f.Log.Fatalf("error creating story folder %v", err)
	}

	extension := f.GetExtension(URL)
	file, err := os.Create(fmt.Sprintf("stories/%s/%d.%s", f.Instagram.Username, time.Now().UnixNano()/1000000, extension))
	if err != nil {
		f.Log.Fatalf("error creating story file %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		f.Log.Fatalf("error copying file %v", err)
	}
}

func NewFile(i *Instagram) *File {
	return &File{
		Instagram: i,
		Log:       log.Default(),
	}
}

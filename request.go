package main

import (
	"log"
	"net/http"
)

type Request struct {
	*File
	*log.Logger
}

func NewRequest() *Request {
	return &Request{
		File:   NewFile(),
		Logger: log.Default(),
	}
}

// Downloading founded url from scraping
func (r *Request) DoRequest(username string, fileURL string, ext string) {
	resp, err := http.Get(fileURL)
	if err != nil {
		r.Logger.Fatal("error when requesting to url: ", err.Error())
	}
	defer resp.Body.Close()

	r.File.ChangeSubFolderName(username)

	err = r.File.Mkdir() // create directory
	if err != nil {
		r.Logger.Fatal("error when make directory: ", err.Error())
	}

	destination, err := r.File.CreateFile(ext) // create stream of created file
	if err != nil {
		r.Logger.Fatal("error when creating a file: ", err.Error())
	}
	defer destination.Close()

	err = r.File.CopyFile(destination, resp.Body) // copy created streamed file to filesystem
	if err != nil {
		r.Logger.Fatal("error copying stream file to filesystem: ", err.Error())
	}
}

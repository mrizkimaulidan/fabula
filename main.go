package main

import (
	"flag"
	"log"
	"sync"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
	"github.com/mrizkimaulidan/fabula/parser"
	"github.com/mrizkimaulidan/fabula/pkg"
)

var (
	username string
)

var (
	flags = []string{"username"}
)

func main() {
	flag.StringVar(&username, "username", "", "the instagram username")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ok, err := pkg.IsFlagPassed(flags)
	if !ok {
		log.Fatal(err.Error())
	}

	instagram := instagram.NewInstagram()
	instagramProfile, err := instagram.GetInstagramProfile(username)
	if err != nil {
		log.Fatal(err.Error())
	}

	fs := file.NewFile(instagramProfile)

	parser := parser.NewParser(instagramProfile, fs)
	response, err := parser.Call()
	if err != nil {
		log.Fatal(err.Error())
	}

	files := parser.Parsing(response)
	log.Printf("found the user with %d story, downloading now please wait..", len(*files))

	err = fs.CreateDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	var wg sync.WaitGroup
	for _, f := range *files {
		wg.Add(1)
		go func(f file.File) {
			defer wg.Done()
			response, err := fs.GetFile(f.URL)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer response.Body.Close()

			createdFile, err := fs.CreateFile(f, response.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer createdFile.Close()

			fs.DownloadText(&f)
		}(f)
	}
	wg.Wait()

	fs.OutputPath()
}

package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
	"github.com/mrizkimaulidan/fabula/parser"
)

func isFlagPassed(args []string) (bool, error) {
	found := false
	i := 0
	flag.Visit(func(f *flag.Flag) {
		if f.Name == args[i] {
			found = true
		} else {
			i++
		}
	})

	return found, fmt.Errorf("missing %s arguments", args[i])
}

func main() {
	var username string
	flag.StringVar(&username, "username", "", "the instagram username")
	flag.Parse()

	ok, err := isFlagPassed([]string{"username"})
	if !ok {
		log.Fatal(err)
	}

	instagram := instagram.NewInstagram()
	instagramProfile, err := instagram.GetInstagramProfile(username)
	if err != nil {
		log.Fatal(err.Error())
	}

	parser := parser.NewParser(*instagramProfile)
	response, err := parser.Call()
	if err != nil {
		log.Fatal(err.Error())
	}

	files := parser.Parsing(response)
	fs := file.NewFile(*instagramProfile)

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
		}(f)
	}
	wg.Wait()
}

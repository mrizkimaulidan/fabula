package main

import (
	"flag"

	"github.com/mrizkimaulidan/fabula/instagram"
	"github.com/mrizkimaulidan/fabula/parser"
)

func main() {
	var profileID string
	flag.StringVar(&profileID, "profileID", "", "the instagram profileID")
	flag.Parse()

	instagram := instagram.Instagram{ProfileID: profileID}
	parser := parser.NewParser(instagram)
	parser.Start()
}

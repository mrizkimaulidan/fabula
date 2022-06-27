package main

import (
	"flag"
	"log"
)

var username string

type Flag struct {
	*log.Logger
}

func NewFlag() *Flag {
	return &Flag{
		Logger: log.Default(),
	}
}

// Checking the flag is passed or not
// returning true when passed
// false when not passed
func (f *Flag) checkIfFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})

	return found
}

// Parse and validate the flags on arguments
// returning map of arguments that provided by user
func (f *Flag) ParsingFlag() map[string]any {
	var args = make(map[string]any)

	flag.StringVar(&username, "username", "", "the instagram username")
	flag.Parse()

	ok := f.checkIfFlagPassed("username")
	if !ok {
		f.Logger.Fatal("missing -username flag, usage: -username=john.doe")
	}

	f.assignArgumentsToMap(args, "username", username)

	return args
}

// Assigning arguments to map
func (f *Flag) assignArgumentsToMap(args map[string]any, name string, value string) {
	args[name] = value
}

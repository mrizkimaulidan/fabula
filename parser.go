package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var URL = "https://igpanda.com/getAjax?type=story&url=%s"

type Parser struct {
	Logger *log.Logger
}

// This function consuming an API from igpanda.
// Unmarshaling the response body to Response struct.
func (p *Parser) Parse(username string) *Response {
	resp, err := http.Get(fmt.Sprintf(URL, username))
	if err != nil {
		p.Logger.Fatalln("error requesting to url", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		p.Logger.Fatalln("error reading response body", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		p.Logger.Fatalln("error unmarshaling response body", err)
	}

	return &response
}

func NewParser() *Parser {
	return &Parser{
		Logger: log.Default(),
	}
}

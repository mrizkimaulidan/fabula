package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://igpanda.com/getAjax?type=story&url=%s"

type Parser struct {
	Log       *log.Logger
	Instagram *Instagram
}

// Call request to igpanda API
func (p *Parser) Call() *Response {
	resp, err := http.Get(fmt.Sprintf(URL, p.Instagram.Username))
	if err != nil {
		p.Log.Fatalf("error requesting to url %v", err)
	}
	defer resp.Body.Close()

	// check status code status
	if resp.StatusCode != 200 {
		p.Log.Fatalf("something went wrong with the API server, %d status code given", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		p.Log.Fatalf("error reading bytes %v", err)
	}

	var body Response
	err = json.Unmarshal(b, &body)
	if err != nil {
		p.Log.Fatalf("error unmarshaling the json %v", err)
	}

	return &Response{
		HTML: body.HTML,
	}
}

func NewParser(i *Instagram) *Parser {
	return &Parser{
		Log:       log.Default(),
		Instagram: i,
	}
}

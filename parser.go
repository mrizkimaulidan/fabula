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
	Instagram *Instagram
}

// Call request to igpanda API
func (r *Parser) Call() *Response {
	resp, err := http.Get(fmt.Sprintf(URL, r.Instagram.Username))
	if err != nil {
		log.Fatalf("error requesting to url %v", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading bytes %v", err)
	}

	var body Response
	err = json.Unmarshal(b, &body)
	if err != nil {
		log.Fatalf("error unmarshaling the json %v", err)
	}

	return &Response{
		HTML: body.HTML,
	}
}

func NewParser(i *Instagram) *Parser {
	return &Parser{
		Instagram: i,
	}
}

package parser

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/mrizkimaulidan/fabula/file"
)

const API_URL = "https://storiesig.info/api/ig/stories/9042373612"

type ParserInterface interface {
	Call() (*Response, error)
	Parsing(response *Response) *[]file.File
}

type Parser struct {
	//
}

func NewParser() ParserInterface {
	return &Parser{}
}

func (p *Parser) Call() (*Response, error) {
	resp, err := http.Get(API_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *Parser) Parsing(response *Response) *[]file.File {
	var files []file.File
	for _, r := range response.Result {
		newFile := file.File{
			Filename: strconv.Itoa(rand.Intn(int(time.Now().UnixNano() / 1000000))),
		}

		if r.HasAudio {
			newFile.Extension = ".mp4"
			newFile.URL = r.VideoVersions[0].URL

			files = append(files, newFile)

		} else {
			newFile.Extension = ".jpg"
			newFile.URL = r.ImageVersions2.Candidates[0].URL

			files = append(files, newFile)
		}
	}

	return &files
}
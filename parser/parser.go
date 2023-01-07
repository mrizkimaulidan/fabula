package parser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
)

const API_URL = "https://storiesig.info/api/ig/stories/%s"

type Parser struct {
	Instagram *instagram.Instagram
	File      *file.File
}

func NewParser(instagram *instagram.Instagram, file *file.File) *Parser {
	return &Parser{
		Instagram: instagram,
		File:      file,
	}
}

// Call the third party API URL to fetching the story
// it will returning the response JSON and decode
// it to Response struct
func (p *Parser) Call() (*Response, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL, p.Instagram.ProfileID))
	if err != nil {
		return nil, fmt.Errorf("error calling request to API URL %s", err.Error())
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body %s", err.Error())
	}

	return &response, nil
}

// Parsing the Response struct and checking
// the story content is video or photo
// returning the slices of struct File
func (p *Parser) Parsing(response *Response) *[]file.File {
	var files []file.File
	for _, r := range response.Result {
		newFile := file.File{
			Filename: p.File.GetRandomString(),
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

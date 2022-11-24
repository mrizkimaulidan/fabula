package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/mrizkimaulidan/fabula/file"
	"github.com/mrizkimaulidan/fabula/instagram"
)

const API_URL = "https://storiesig.info/api/ig/stories/%s"

type ParserInterface interface {
	Call() (*Response, error)
	Parsing(response *Response) *[]file.File
	Start()
}

type Parser struct {
	InstagramProfile instagram.InstagramProfile
}

func NewParser(instagram instagram.InstagramProfile) ParserInterface {
	return &Parser{
		InstagramProfile: instagram,
	}
}

func (p *Parser) Call() (*Response, error) {
	resp, err := http.Get(fmt.Sprintf(API_URL, p.InstagramProfile.Users[0].User.Pk))
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

func (p *Parser) Start() {
	resp, err := p.Call()
	if err != nil {
		log.Fatal(err.Error())
	}

	fileS := file.NewFile()
	fileS.CreateDir()

	files := p.Parsing(resp)
	for _, f := range *files {
		// TODO: goroutine??
		resp, err := fileS.GetFile(f.URL)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer resp.Body.Close()

		createdFile, err := fileS.CreateFile(f, resp.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer createdFile.Close()
	}
}

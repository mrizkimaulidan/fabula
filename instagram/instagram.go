package instagram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InstagramInterface interface {
	GetProfileIDByUsername(username string) (*InstagramProfile, error)
}

type Instagram struct {
	ProfileID string
}

func NewInstagram() InstagramInterface {
	return &Instagram{}
}

func (i *Instagram) GetProfileIDByUsername(username string) (*InstagramProfile, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/web/search/topsearch/?query=%s", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response InstagramProfile
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

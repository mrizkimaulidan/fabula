package main

type Highlight struct {
	Result []struct {
		ImageVersions2 struct {
			Candidates []struct {
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				URL          string `json:"url"`
				URLSignature struct {
					Expires   int    `json:"expires"`
					Signature string `json:"signature"`
				} `json:"url_signature"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		OriginalHeight int    `json:"original_height"`
		OriginalWidth  int    `json:"original_width"`
		Pk             string `json:"pk"`
		TakenAt        int    `json:"taken_at"`
		VideoVersions  []struct {
			Height       int    `json:"height"`
			Type         int    `json:"type"`
			URL          string `json:"url"`
			Width        int    `json:"width"`
			URLSignature struct {
				Expires   int    `json:"expires"`
				Signature string `json:"signature"`
			} `json:"url_signature"`
		} `json:"video_versions,omitempty"`
		HasAudio bool `json:"has_audio,omitempty"`
	} `json:"result"`
}

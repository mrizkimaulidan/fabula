package main

type HighlightList struct {
	Result []struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		CoverMedia struct {
			CroppedImageVersion struct {
				URL          string `json:"url"`
				URLSignature struct {
					Expires   string `json:"expires"`
					Signature string `json:"signature"`
				} `json:"url_signature"`
			} `json:"cropped_image_version"`
		} `json:"cover_media"`
	} `json:"result"`
}

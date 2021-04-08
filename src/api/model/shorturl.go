package model

type ShorturlRequest struct {
	Name     string `json:"name"`
	Shorturl string `json:"short_url"`
	URL      string `json:"url"`
}

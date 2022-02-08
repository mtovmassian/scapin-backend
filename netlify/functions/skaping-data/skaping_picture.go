package main

type SkapingPicture struct {
	Url string `json:"url"`
}

func NewSkapingPicture(url string) *SkapingPicture {
	return &SkapingPicture{Url: url}
}

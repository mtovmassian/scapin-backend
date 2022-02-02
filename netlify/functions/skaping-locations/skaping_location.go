package main

import (
	"encoding/json"
)

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type SkapingLocation struct {
	Url      string `json:"url"`
	Position LatLng `json:"position"`
	Title    string `json:"title"`
}

func NewSkapingLocation(url string, position LatLng, title string) *SkapingLocation {
	return &SkapingLocation{Url: url, Position: position, Title: title}
}

func (skapingLocation *SkapingLocation) toJson() string {
	jsonData, _ := json.Marshal(skapingLocation)
	return string(jsonData)
}

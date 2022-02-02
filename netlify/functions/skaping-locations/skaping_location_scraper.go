package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type SkapingLocationScraper struct {
	rawHtml string
}

func NewSkapingLocationScraper(rawHtml string) *SkapingLocationScraper {
	return &SkapingLocationScraper{rawHtml: rawHtml}
}

func NewSkapingLocationScraperFromUrl(url string) *SkapingLocationScraper {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return &SkapingLocationScraper{rawHtml: string(body)}
}

func (scraper *SkapingLocationScraper) ScrapLocations() SkapingLocations {
	skapingLocations := []SkapingLocation{}
	for _, rawDataLocation := range scraper.ScrapRawDataLocations() {
		skapingLocations = append(skapingLocations, scraper.FromRawToSkapingLocation(&rawDataLocation))
	}

	return skapingLocations
}

func (scraper *SkapingLocationScraper) ScrapRawDataLocations() []string {
	pattern := regexp.MustCompile(`(new google\.maps\.InfoWindow[.\s\S]*?new google\.maps\.Marker[.\s\S]*?;)`)
	matches := pattern.FindAllStringSubmatch(scraper.rawHtml, -1)
	rawDataLocations := []string{}
	for _, match := range matches {
		rawDataLocations = append(rawDataLocations, match[1])
	}
	return rawDataLocations
}

func (scraper *SkapingLocationScraper) ExractLocationUrl(rawDataLocation *string) string {
	pattern := regexp.MustCompile(`href=\\"(.*)\\"`)
	matches := pattern.FindAllStringSubmatch(*rawDataLocation, -1)
	if len(matches) == 0 {
		return ""
	}

	return matches[0][1]
}

func (scraper *SkapingLocationScraper) ExtractLocationLatLng(rawDataLocation *string) LatLng {
	pattern := regexp.MustCompile(`position: new google.maps.LatLng\((.*),\s?(.*)\)`)
	matches := pattern.FindAllStringSubmatch(*rawDataLocation, -1)
	if len(matches) == 0 {
		return LatLng{}
	}
	lat, _ := strconv.ParseFloat(matches[0][1], 64)
	lng, _ := strconv.ParseFloat(matches[0][2], 64)

	return LatLng{lat, lng}
}

func (scraper *SkapingLocationScraper) ExtractLocationTitle(rawDataLocation *string) string {
	pattern := regexp.MustCompile(`title:"(.*)"`)
	matches := pattern.FindAllStringSubmatch(*rawDataLocation, -1)
	if len(matches) == 0 {
		return ""
	}

	return matches[0][1]
}

func (scraper *SkapingLocationScraper) FromRawToSkapingLocation(rawDataLocation *string) SkapingLocation {
	return SkapingLocation{
		Url:      scraper.ExractLocationUrl(rawDataLocation),
		Position: scraper.ExtractLocationLatLng(rawDataLocation),
		Title:    scraper.ExtractLocationTitle(rawDataLocation),
	}
}

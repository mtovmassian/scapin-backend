package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

type SkapingPictureScraper struct {
	rawHtml string
}

func NewSkapingPictureScraper(rawHtml string) *SkapingPictureScraper {
	return &SkapingPictureScraper{rawHtml: rawHtml}
}

func NewSkapingPictureScraperFromUrl(url string) (scraper *SkapingPictureScraper, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &SkapingPictureScraper{rawHtml: string(body)}, nil
}

func (scraper *SkapingPictureScraper) ScrapPicture() SkapingPicture {
	return SkapingPicture{
		Url: scraper.ExractPictureUrl(),
	}
}

func (scraper *SkapingPictureScraper) ExractPictureUrl() string {
	pattern := regexp.MustCompile(`<link[.\s]*rel="image_src"[.\s]*href="(.*)"[.\s]*/>`)
	matches := pattern.FindAllStringSubmatch(scraper.rawHtml, -1)
	if len(matches) == 0 {
		return ""
	}

	return matches[0][1]
}

// func (scraper *SkapingPictureScraper) ScrapRawDataPictures() []string {
// 	pattern := regexp.MustCompile(`(new google\.maps\.InfoWindow[.\s\S]*?new google\.maps\.Marker[.\s\S]*?;)`)
// 	matches := pattern.FindAllStringSubmatch(scraper.rawHtml, -1)
// 	rawDataPictures := []string{}
// 	for _, match := range matches {
// 		rawDataPictures = append(rawDataPictures, match[1])
// 	}
// 	return rawDataPictures
// }

// func (scraper *SkapingPictureScraper) ExractPictureUrl(rawDataPicture *string) string {
// 	pattern := regexp.MustCompile(`href=\\"(.*)\\"`)
// 	matches := pattern.FindAllStringSubmatch(*rawDataPicture, -1)
// 	if len(matches) == 0 {
// 		return ""
// 	}

// 	return matches[0][1]
// }

// func (scraper *SkapingPictureScraper) ExtractPictureLatLng(rawDataPicture *string) LatLng {
// 	pattern := regexp.MustCompile(`position: new google.maps.LatLng\((.*),\s?(.*)\)`)
// 	matches := pattern.FindAllStringSubmatch(*rawDataPicture, -1)
// 	if len(matches) == 0 {
// 		return LatLng{}
// 	}
// 	lat, _ := strconv.ParseFloat(matches[0][1], 64)
// 	lng, _ := strconv.ParseFloat(matches[0][2], 64)

// 	return LatLng{lat, lng}
// }

// func (scraper *SkapingPictureScraper) ExtractPictureTitle(rawDataPicture *string) string {
// 	pattern := regexp.MustCompile(`title:"(.*)"`)
// 	matches := pattern.FindAllStringSubmatch(*rawDataPicture, -1)
// 	if len(matches) == 0 {
// 		return ""
// 	}

// 	return matches[0][1]
// }

// func (scraper *SkapingPictureScraper) FromRawToSkapingPicture(rawDataPicture *string) SkapingPicture {
// 	return SkapingPicture{
// 		Url:      scraper.ExractPictureUrl(rawDataPicture),
// 		Position: scraper.ExtractPictureLatLng(rawDataPicture),
// 		Title:    scraper.ExtractPictureTitle(rawDataPicture),
// 	}
// }

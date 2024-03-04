package main

import (
	"io"
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

	body, err := io.ReadAll(resp.Body)
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

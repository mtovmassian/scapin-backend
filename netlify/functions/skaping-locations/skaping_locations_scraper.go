package main

import (
	"regexp"
)

type SkapingRawDataLocation = string

type SkapingLocation struct {
}

func ScrapSkapingRawDataLocations(rawHtml *string) []SkapingRawDataLocation {
	pattern := regexp.MustCompile(`(new google\.maps\.InfoWindow[.\s\S]*?new google\.maps\.Marker[.\s\S]*?;)`)
	matches := pattern.FindAllStringSubmatch(*rawHtml, -1)
	rawDataLocations := []string{}
	for _, match := range matches {
		rawDataLocations = append(rawDataLocations, match[1])
	}
	return rawDataLocations
}

func ExractSkapingLocationUrl(rawDataLocation *SkapingRawDataLocation) string {
	pattern := regexp.MustCompile(`href=\\"(.*)\\"`)
	matches := pattern.FindAllStringSubmatch(*rawDataLocation, -1)
	if len(matches) == 0 {
		return ""
	}

	return matches[0][1]
}

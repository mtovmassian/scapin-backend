package main

import (
	"regexp"
	"strconv"
)

type LatLng struct {
	Lat float64
	Lng float64
}

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

func ExtractSkapingLocationLatLng(rawDataLocation *SkapingRawDataLocation) LatLng {
	pattern := regexp.MustCompile(`position: new google.maps.LatLng\((.*),\s?(.*)\)`)
	matches := pattern.FindAllStringSubmatch(*rawDataLocation, -1)
	if len(matches) == 0 {
		return LatLng{}
	}
	lat, _ := strconv.ParseFloat(matches[0][1], 64)
	lng, _ := strconv.ParseFloat(matches[0][2], 64)

	return LatLng{lat, lng}
}

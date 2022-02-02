package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJson_NominalCase_String(t *testing.T) {
	// GIVEN
	skapingLocation := NewSkapingLocation(
		"uniform",
		LatLng{1.1, 2.2},
		"tango",
	)
	// WHEN
	skapingLocationJson := skapingLocation.ToJson()
	expectedSkapingLocationJson := strings.Join([]string{
		"{",
		"\"url\":\"uniform\",",
		"\"position\":{\"lat\":1.1,\"lng\":2.2},",
		"\"title\":\"tango\"",
		"}",
	}, "")
	// THEN
	assert.Equal(t, expectedSkapingLocationJson, skapingLocationJson)
}

func TestToJsonAll_NominalCase_String(t *testing.T) {
	// GIVEN
	skapingLocation1 := NewSkapingLocation(
		"uniform",
		LatLng{1.1, 2.2},
		"tango",
	)
	skapingLocation2 := NewSkapingLocation(
		"victor",
		LatLng{3.3, 4.4},
		"x-ray",
	)
	// WHEN
	var skapingLocations SkapingLocations = []SkapingLocation{*skapingLocation1, *skapingLocation2}
	skapingLocationsJson := skapingLocations.ToJson()
	expectedSkapingLocationsJson := strings.Join([]string{
		"[",
		"{",
		"\"url\":\"uniform\",",
		"\"position\":{\"lat\":1.1,\"lng\":2.2},",
		"\"title\":\"tango\"",
		"},",
		"{",
		"\"url\":\"victor\",",
		"\"position\":{\"lat\":3.3,\"lng\":4.4},",
		"\"title\":\"x-ray\"",
		"}",
		"]",
	}, "")
	// THEN
	assert.Equal(t, expectedSkapingLocationsJson, skapingLocationsJson)
}

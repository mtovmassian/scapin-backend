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
	skapingLocationJson := skapingLocation.toJson()
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

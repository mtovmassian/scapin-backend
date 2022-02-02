package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSkapingLocationScraperFromUrl(t *testing.T) {
	t.Skip("Skipping integration test.")
	// GIVEN
	url := "https://www.skaping.com/camera/map"
	// WHEN
	scraper := NewSkapingLocationScraperFromUrl(url)
	// THEN
	fmt.Printf("%v", scraper.ScrapLocations())
	assert.NotEmpty(t, scraper.rawHtml)
}

func TestExtractLocationUrl_NominalCase_String(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	rawDataLocationNoMatch := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" ferh=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	locationUrl := NewSkapingLocationScraper("").ExractLocationUrl(&rawDataLocation)
	locationUrlNoMatch := NewSkapingLocationScraper("").ExractLocationUrl(&rawDataLocationNoMatch)
	// THEN
	expectedLocationUrl := "http://www.skaping.com/banffgondola"
	assert.Equal(t, expectedLocationUrl, locationUrl)
	assert.Empty(t, locationUrlNoMatch)
}

func TestExractLatLng_NominalCase_LatLng(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	rawDataLocationInvalidFloats := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(alpha, bravo),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	latLng := NewSkapingLocationScraper("").ExtractLocationLatLng(&rawDataLocation)
	latLngInvalidFloats := NewSkapingLocationScraper("").ExtractLocationLatLng(&rawDataLocationInvalidFloats)
	// THEN
	assert.Equal(t, LatLng{51.14460700, -115.57476600}, latLng)
	assert.Equal(t, LatLng{0, 0}, latLngInvalidFloats)
}

func TestExractLocationTitle_NominalCase_String(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	locationTitle := NewSkapingLocationScraper("").ExtractLocationTitle(&rawDataLocation)
	// THEN
	expectedLocationTitle := "Banff Gondola"
	assert.Equal(t, expectedLocationTitle, locationTitle)
}

func TestFromRawToSkapingLocation_NominalCase_SkapingLocation(t *testing.T) {
	// GIVEN
	rawDataLocation := `new google.maps.InfoWindow({
		content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/banffgondola\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/banffgondola\">http://www.skaping.com/banffgondola</a>"
	});
	 markers[22] = new google.maps.Marker({
		 position: new google.maps.LatLng(51.14460700, -115.57476600),
		 animation: google.maps.Animation.DROP,
		 map: map,
		 title:"Banff Gondola"
	 });`
	// WHEN
	skapingLocation := NewSkapingLocationScraper("").FromRawToSkapingLocation(&rawDataLocation)
	// THEN
	expectedSkapingLocation := SkapingLocation{
		Url:      "http://www.skaping.com/banffgondola",
		Position: LatLng{51.14460700, -115.57476600},
		Title:    "Banff Gondola",
	}
	assert.Equal(t, expectedSkapingLocation, skapingLocation)
}

func TestScrapRawDataLocations_NominalCase_ListOfRawDataLocations(t *testing.T) {
	// GIVEN
	rawHtml := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			windows[1] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/pra-loup/molanes\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/pra-loup/molanes\">http://www.skaping.com/pra-loup/molanes</a>"
			});
			markers[1] = new google.maps.Marker({
				position: new google.maps.LatLng(44.36097000, 6.60417300),
				animation: google.maps.Animation.DROP,
				map: map,
				title:" PRA LOUP - CLOS DU SERRE (1820m)"
			});
			markers[1].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[1].open(map, markers[1]);
				openedWindow = windows[1];
			});
			bounds.extend(markers[1].getPosition());
		

			windows[2] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/ski-macedonia/bistra-mavrovo\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/ski-macedonia/bistra-mavrovo\">http://www.skaping.com/ski-macedonia/bistra-mavrovo</a>"
			});
			markers[2] = new google.maps.Marker({
				position: new google.maps.LatLng(41.64767700, 20.73542600),
				animation: google.maps.Animation.DROP,
				map: map,
				title:"#skimacedonia Bistra-Mavrovo"
			});
			markers[2].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[2].open(map, markers[2]);
				openedWindow = windows[2];
			});
			bounds.extend(markers[2].getPosition());
		</script>
	</head>
	</html>`
	rawHtmlNoMatch := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			console.log('Charly Lima');
		</script>
	</head>
	</html>
	`
	// WHEN
	rawDataLocations := NewSkapingLocationScraper(rawHtml).ScrapRawDataLocations()
	rawDataLocationsNoMatch := NewSkapingLocationScraper(rawHtmlNoMatch).ScrapRawDataLocations()
	// THEN
	expectedRawDataLocation1 := `new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/pra-loup/molanes\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/pra-loup/molanes\">http://www.skaping.com/pra-loup/molanes</a>"
			});
			markers[1] = new google.maps.Marker({
				position: new google.maps.LatLng(44.36097000, 6.60417300),
				animation: google.maps.Animation.DROP,
				map: map,
				title:" PRA LOUP - CLOS DU SERRE (1820m)"
			});`
	expectedRawDataLocation2 := `new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/ski-macedonia/bistra-mavrovo\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/ski-macedonia/bistra-mavrovo\">http://www.skaping.com/ski-macedonia/bistra-mavrovo</a>"
			});
			markers[2] = new google.maps.Marker({
				position: new google.maps.LatLng(41.64767700, 20.73542600),
				animation: google.maps.Animation.DROP,
				map: map,
				title:"#skimacedonia Bistra-Mavrovo"
			});`
	expectedRawDataLocations := []string{
		expectedRawDataLocation1,
		expectedRawDataLocation2,
	}
	assert.Equal(t, expectedRawDataLocations, rawDataLocations)
	assert.Empty(t, rawDataLocationsNoMatch)
}

func TestScrapSkapingLocations_NominalCase_ListOfSkapingLocations(t *testing.T) {
	// GIVEN
	rawHtml := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			windows[1] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/pra-loup/molanes\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/pra-loup/molanes\">http://www.skaping.com/pra-loup/molanes</a>"
			});
			markers[1] = new google.maps.Marker({
				position: new google.maps.LatLng(44.36097000, 6.60417300),
				animation: google.maps.Animation.DROP,
				map: map,
				title:" PRA LOUP - CLOS DU SERRE (1820m)"
			});
			markers[1].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[1].open(map, markers[1]);
				openedWindow = windows[1];
			});
			bounds.extend(markers[1].getPosition());
		

			windows[2] = new google.maps.InfoWindow({
				content: "<div style=\"width:650px;max-width:100%\"><div id=\"skaping\" style=\"position:relative;width:100%;padding-top:56.25%;\"><iframe src=\"//www.skaping.com/ski-macedonia/bistra-mavrovo\" allowfullscreen style=\"position:absolute;top:0;left:0;height:100%;width:100%;border:0\"></iframe></div></div><br /><a target=\"_blank\" href=\"http://www.skaping.com/ski-macedonia/bistra-mavrovo\">http://www.skaping.com/ski-macedonia/bistra-mavrovo</a>"
			});
			markers[2] = new google.maps.Marker({
				position: new google.maps.LatLng(41.64767700, 20.73542600),
				animation: google.maps.Animation.DROP,
				map: map,
				title:"#skimacedonia Bistra-Mavrovo"
			});
			markers[2].addListener('click', function() {
				if (openedWindow) {
					openedWindow.close();
				}
				windows[2].open(map, markers[2]);
				openedWindow = windows[2];
			});
			bounds.extend(markers[2].getPosition());
		</script>
	</head>
	</html>`
	rawHtmlNoMatch := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<script>
			console.log('Charly Lima');
		</script>
	</head>
	</html>
	`
	// WHEN
	skapingLocations := NewSkapingLocationScraper(rawHtml).ScrapLocations()
	skapingLocationsNoMatch := NewSkapingLocationScraper(rawHtmlNoMatch).ScrapLocations()
	// THEN
	expectedSkapingLocation1 := SkapingLocation{
		Url:      "http://www.skaping.com/pra-loup/molanes",
		Position: LatLng{44.36097000, 6.60417300},
		Title:    " PRA LOUP - CLOS DU SERRE (1820m)",
	}
	expectedSkapingLocation2 := SkapingLocation{
		Url:      "http://www.skaping.com/ski-macedonia/bistra-mavrovo",
		Position: LatLng{41.64767700, 20.73542600},
		Title:    "#skimacedonia Bistra-Mavrovo",
	}
	expectedSkapingLocations := []SkapingLocation{
		expectedSkapingLocation1,
		expectedSkapingLocation2,
	}
	assert.Equal(t, expectedSkapingLocations, skapingLocations)
	assert.Empty(t, skapingLocationsNoMatch)
}

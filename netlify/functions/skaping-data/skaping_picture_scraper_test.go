package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSkapingPictureScraperFromUrl(t *testing.T) {
	t.Skip("Skipping integration test.")
	// GIVEN
	url := "http://www.skaping.com/alpedhuez/herpie"
	// WHEN
	scraper, _ := NewSkapingPictureScraperFromUrl(url)
	// THEN
	fmt.Printf("%v", scraper.ScrapPicture())
	assert.NotEmpty(t, scraper.rawHtml)
}

func TestExractPictureUrl_NominalCase_String(t *testing.T) {
	// GIVEN
	rawHtml := `<!DOCTYPE html>
	<html class="no-js">
	<head>
		<meta charset="utf-8" /><title>Alpe d'Huez  - Herpie</title>
		<meta name="description" content="Webcam à l'Alpe d'Huez - Herpie" />
		<meta name="keywords" content="alpe d'huez webcam livecam landscape panoramique live" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no, height=device-height, target-densitydpi=device-dpi">
		<link rel="preload" href="/fonts/Muli-Regular.woff2" as="font" type="font/woff2" crossorigin="anonymous">
		<link rel="preload" href="/fonts/Muli-Bold.woff2" as="font" type="font/woff2" crossorigin="anonymous">
		<link rel="preload" href="/fonts/Muli-Italic.woff2" as="font" type="font/woff2" crossorigin="anonymous">
		<link rel="preload" href="/fonts/Muli-Light.woff2" as="font" type="font/woff2" crossorigin="anonymous">
		<link rel="apple-touch-icon" sizes="57x57" href="/apple-touch-icon-57x57.png?v=PYeJgwJrQP">
		<link rel="apple-touch-icon" sizes="60x60" href="/apple-touch-icon-60x60.png?v=PYeJgwJrQP">
		<link rel="apple-touch-icon" sizes="72x72" href="/apple-touch-icon-72x72.png?v=PYeJgwJrQP">
		<link rel="apple-touch-icon" sizes="76x76" href="/apple-touch-icon-76x76.png?v=PYeJgwJrQP">
		<link rel="apple-touch-icon" sizes="114x114" href="/apple-touch-icon-114x114.png?v=PYeJgwJrQP">
		<link rel="apple-touch-icon" sizes="120x120" href="/apple-touch-icon-120x120.png?v=PYeJgwJrQP">
		<link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png?v=PYeJgwJrQP">
		<link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png?v=PYeJgwJrQP">
		<link rel="manifest" href="/site.webmanifest?v=PYeJgwJrQP">
		<link rel="mask-icon" href="/safari-pinned-tab.svg?v=PYeJgwJrQP" color="#333333">
		<link rel="shortcut icon" href="/favicon.ico?v=PYeJgwJrQP">
		<meta name="msapplication-TileColor" content="#333333">
		<meta name="HandheldFriendly" content="true" />
		<meta name="mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-status-bar-style" content="black">
		<link rel="image_src" href="http://data.skaping.com/alpe-dhuez/herpie/2022/02/08/large/09-20.jpg" />
		<link rel="canonical" href="https://www.skaping.com/alpedhuez/herpie" />
	</head>
	<body>
	</body`
	// WHEN
	pictureUrl := NewSkapingPictureScraper(rawHtml).ExractPictureUrl()
	pictureUrlNoMatch := NewSkapingPictureScraper("").ExractPictureUrl()
	// THEN
	expectedPictureUrl := "http://data.skaping.com/alpe-dhuez/herpie/2022/02/08/large/09-20.jpg"
	assert.Equal(t, expectedPictureUrl, pictureUrl)
	assert.Empty(t, pictureUrlNoMatch)
}

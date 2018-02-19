package filterurl

import (
	"net/url"
	"strings"
)

/*
Filter ...
*/
func Filter(urls []string) []string {
	var validurls []string
	for _, rawurl := range urls {
		if IsValidURL(rawurl) {
			validurls = append(validurls, strings.Fields(rawurl)[0])
		}
	}
	return validurls
}

/*
IsValidURL ...
*/
func IsValidURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	return true
}

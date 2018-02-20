package filter

import (
	"net/url"
	"strings"
)

/*
GetValidURLs ... This method filters all valid URLs from given list of URLs
*/
func GetValidURLs(urls []string) []string {
	var validurls []string
	for _, rawurl := range urls {
		if IsValidURL(rawurl) {
			validurls = append(validurls, strings.Fields(rawurl)[0])
		}
	}
	return validurls
}

/*
IsValidURL ... This method checks if given URL is a valid URL
*/
func IsValidURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}
	return true
}

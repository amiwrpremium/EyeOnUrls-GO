package utils

import "regexp"

// IsValidURL is a function that checks if a given URL is valid
// @param url string
// @return bool
func IsValidURL(url string) bool {
	regex := regexp.MustCompile(`^((http|https)://)?[a-z0-9]+([-.][a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(/.*)?$`)
	if !regex.MatchString(url) {
		return false
	}
	return true
}

// CompleteURL is a function that completes a URL with the http:// prefix if it's missing
// @param url string
// @return string
func CompleteURL(url string) string {
	regex := regexp.MustCompile(`^(http|https)://`)
	if !regex.MatchString(url) {
		return "http://" + url // pragma: no cover
	}
	return url
}

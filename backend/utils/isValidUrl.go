package utils

import "net/url"

func IsValidUrl(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return true
}
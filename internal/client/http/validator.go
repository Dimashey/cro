package http

import "net/url"

func isValidUrl(URL string) bool {
 	u, err := url.ParseRequestURI(URL);
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}

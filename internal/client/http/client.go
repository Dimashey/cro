package http

import "fmt"
import "strings"
import "net/http" 

type client struct {
	baseUrl string
	http    *http.Client
}

func New(baseURL string) *client {
	isValid := isValidUrl(baseURL)

	if !isValid {
		panic(fmt.Sprintf("invalid url was provided: %s\n", baseURL))
	}

	return &client{baseUrl: strings.TrimSuffix(baseURL, "/"), http: &http.Client{}}
}

func (c client) Get(endpoint string) (*http.Response, error) {
	path := c.path(endpoint)

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (c client) path(endpoint string) string {
	path := c.baseUrl + "/" + strings.TrimPrefix(endpoint, "/")

	return path
}


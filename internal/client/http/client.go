package http

import "strings"
import "net/http" 
import "errors"

type Client struct {
	baseUrl string
	http    *http.Client
}

var InvalidUrl = errors.New("Invalid URL")

func New(baseURL string) (*Client, error) {
	isValid := isValidUrl(baseURL)

	if !isValid {
		return nil, InvalidUrl
	}

	return &Client{baseUrl: strings.TrimSuffix(baseURL, "/"), http: &http.Client{}}, nil
}

func (c Client) Get(endpoint string, query map[string]string) (*http.Response, error) {
	path := c.path(endpoint)

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	if query != nil {
		q := req.URL.Query()

		for key, value := range query {
			q.Add(key, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}


func (c Client) path(endpoint string) string {
	path := c.baseUrl + "/" + strings.TrimPrefix(endpoint, "/")

	return path
}


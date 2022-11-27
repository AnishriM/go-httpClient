package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, header http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	request.Header = c.getFullHeaders(header)
	return client.Do(request)
}

func (c *httpClient) getFullHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// Add common headers
	for key, value := range c.header {
		if len(value) > 0 {
			result.Set(key, value[0])
		}
	}

	// Add custom headers
	for key, value := range headers {
		if len(value[0]) > 0 {
			result.Set(key, value[0])
		}
	}

	return result
}

package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string, url string, header http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	fullHeader := c.getFullHeaders(header)
	fullBody, err := c.getRequestBody(fullHeader.Get("contentType"), body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(fullBody))
	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	request.Header = fullHeader
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

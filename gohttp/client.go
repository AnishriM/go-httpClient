package gohttp

import "net/http"

type httpClient struct {
	header http.Header
}

func New() HTTPClient {
	return &httpClient{}
}

type HTTPClient interface {
	SetHeaders(header http.Header)
	GET(url string, header http.Header) (*http.Response, error)
	POST(url string, header http.Header, body interface{}) (*http.Response, error)
	PUT(url string, header http.Header, body interface{}) (*http.Response, error)
	PATCH(url string, header http.Header, body interface{}) (*http.Response, error)
	DELETE(url string, header http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(header http.Header) {
	c.header = header
}

func (c *httpClient) GET(url string, header http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, header, nil)
}
func (c *httpClient) POST(url string, header http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, header, body)
}
func (c *httpClient) PUT(url string, header http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, header, body)
}
func (c *httpClient) PATCH(url string, header http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, header, body)
}
func (c *httpClient) DELETE(url string, header http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, header, nil)
}

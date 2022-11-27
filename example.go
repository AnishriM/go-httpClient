package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AnishriM/go-httpClient/gohttp"
)

var (
	client = gitHubHttpClient()
)

func gitHubHttpClient() gohttp.HTTPClient {
	client := gohttp.New()
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)
	return client
}

func main() {
	response, err := client.GET("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

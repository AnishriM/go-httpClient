package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AnishriM/go-httpClient/gohttp"
)

func main() {
	client := gohttp.New()
	response, err := client.GET("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

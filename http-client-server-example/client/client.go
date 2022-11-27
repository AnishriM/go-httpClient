/*
@Author: Sagar Mhantati

  - In below code we created http GET request which send request to localhost and on 3001 port.
  - As server is listening on same port this request will get fulfiled by server.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpMethod := "GET"
	url := "http://localhost:3001/time"

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/xml")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println("Status Code: ", response.StatusCode, "Body:", string(bytes))
}

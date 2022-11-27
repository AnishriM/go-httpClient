/*
@Author: Sagar Mhantati

	- In below code we have created mux which is basically map which stores information about <url, handler>. We can register multiple <url, handler> with mux.
	- After that we used ListenAndServer function to listen on 3001 port and fulfil those requests.
*/

package main

import (
	"net/http"
)

func main() {
	//Here mux maps multiple handler with url.
	mux := http.NewServeMux()

	//Handler is controller it process and send response
	rh := http.RedirectHandler("google.com", 307)

	//Here we adding map /time url with rh handler so that reuqest which comes
	//at /time url will be fulfilled by rh handler
	mux.Handle("/time", rh)

	//3001 port is opened and served for http requests.
	http.ListenAndServe(":3001", mux)
}

/*
@Author: Sagar Mhantati
In below code we have created mux which is basically map which stores information about
<url, handler>. We can register multiple <url, handler> with mux.
After that we created timeHandler struct and created function "ServeHTTP" which will serve
http requests. Here we have created our own handler.
while registring mux, mux check ServeHTTP function is implemented or not? As our structure
implement this function so its fulfil mux's requirement.
After that we used ListenAndServer function to listen on 3001 port and fulfil those requests.

In sample-1.go we used RedirectHandler which provided by net/http package. And in sample-2.go
we created our own handler.
*/
package main

import (
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (t timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(t.format)
	w.Write([]byte(currentTime))
}

func main() {
	mux := http.NewServeMux()
	th := timeHandler{time.RFC1123}
	mux.Handle("/time", th)
	http.ListenAndServe(":3002", mux)
}

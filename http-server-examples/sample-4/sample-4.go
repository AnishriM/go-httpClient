/*
@Author: Sagar Mhantati

  - In below code we have created mux which is basically map which stores information about
	<url, handler>. We can register multiple <url, handler> with mux. (Line:40)

  - After that we created 'timeHandler' function which take foramt as a input. Inside this function
    we implemented signature of "ServeHTTP" and returned this function as a return value. (Line: 32 function block)
	Here signature is same as that "ServeHTTP" but function is annonyomous function so to work this annonymous function
	as a valid handler we used http.HandlerFunc (and it will internally use 'ServeHTTP' function).

  - After that we used ListenAndServer function to listen on 3002 port and fulfil those requests.

Conclusion:
  - In sample-1.go we used RedirectHandler which provided by net/http package.
  - In sample-2.go we created our own handler using struct.
  - In sample-3.go we created own handler function without creating struct.
  - In sample-4.go we created own handler function without creating struct. In sample-3.go we
    used static format but now in sample-4.go we are passing format as input. So format can be changed
	dynamically according to provided input.
*/

package main

import (
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	time := time.Now().Format(format)
	timehandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time))
	}
	return http.HandlerFunc(timehandler)
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/time", timeHandler(time.RFC1123))
	http.ListenAndServe(":3002", mux)
}

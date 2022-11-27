/*
@Author: Sagar Mhantati

  - In below code we have not but we have used default mux which is provided by http package.
    Here we are not exposing mux and which prevent from third party attacks. (Line:38)

  - After that we created 'timeHandler' function which take foramt as a input. Inside this function
    we implemented signature of "ServeHTTP" and returned this function as a return value. (Line: 32 function block)
	Here signature is same as that "ServeHTTP" but function is annonyomous function so to work this annonymous function
	as a valid handler we used http.HandlerFunc (and it will internally use 'ServeHTTP' function).

  - After that we used ListenAndServer function to listen on 3004 port and fulfil those requests.

Conclusion:
  - In sample-1.go we used RedirectHandler which provided by net/http package.
  - In sample-2.go we created our own handler using struct.
  - In sample-3.go we created own handler function without creating struct.
  - In sample-4.go we created own handler function without creating struct. In sample-3.go we
    used static format but now in sample-4.go we are passing format as input. So format can be changed
	dynamically according to provided input.
  - In sample-5.go almost all things are same only difference is that in sample-5.go we used default mux 
    which is provided by http package.
*/

package main

import (
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	th := func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().Format(format)
		w.Write([]byte(time))
	}
	return http.HandlerFunc(th)
}
func main() {
	http.Handle("/time", timeHandler(time.RFC1123))
	http.ListenAndServe(":3004", nil)
}

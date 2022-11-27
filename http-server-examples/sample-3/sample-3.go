/*
@Author: Sagar Mhantati

  - In below code we have created mux which is basically map which stores information about
	<url, handler>. We can register multiple <url, handler> with mux.

  - After that we created 'timeHandler' function which has signature same as that of "ServeHTTP"
	and this function will serve http requests. Here we have created our own handler function.

  - while registring mux, mux check ServeHTTP function is implemented or not? In our handler
	function.

  - we have used same signature as that of ServeHTTP but function name is different. So in this case
	http.HandlerFunc comes into picture. HandlerFunc will validate 'timeHandler' and it will internally
	use 'ServeHTTP' function.

  - After that we used ListenAndServer function to listen on 3001 port and fulfil those requests.

Conclusion: 
  - In sample-1.go we used RedirectHandler which provided by net/http package.
  - In sample-2.go we created our own handler using struct.
  - In sample-3.go we created own handler function without creating struct.
*/

package main

import (
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	time := time.Now().Format(time.RFC1123)
	w.Write([]byte(time))
}

func main() {
	mux := http.NewServeMux()
	th := http.HandlerFunc(timeHandler)
	mux.Handle("/time", th)

	http.ListenAndServe(":3001", mux)
}

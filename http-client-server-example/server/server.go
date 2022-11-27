/*
@Author: Sagar Mhantati

  - In below code we created http server which listen and serve on 3001 port.
*/
package main

import "net/http"

func cutomHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/time", cutomHandler)
	http.ListenAndServe(":3001", nil)
}

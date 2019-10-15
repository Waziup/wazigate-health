package main

import "net/http"

var www = http.FileServer(http.Dir("www"))

// Serve is the HTTP Handler
func Serve(resp http.ResponseWriter, req *http.Request) {

	www.ServeHTTP(resp, req)
}

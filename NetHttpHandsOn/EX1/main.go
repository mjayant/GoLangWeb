/*
Problem statement :
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name
*/

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/me", meHandler)
	http.ListenAndServe(":8000", nil)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is default handler")
}

func dogHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is dog handler")
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is me hadler")
	log.Println("request object", req)
	url_query := req.URL.Query()
	name := url_query.Get("name")
	log.Println("Passed name param value", name)
	header_val := req.Header.Get("X-Custom-Header")
	log.Println("Header value ", header_val)
}

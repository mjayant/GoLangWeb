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
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/me", meHandler)
	http.ListenAndServe(":8000", nil)
}

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	url_query := req.URL.Query()
	name := url_query.Get("name")
	data := data{
		DataType: "Default",
		Data:     name,
	}
	tpl.Execute(res, data)
}

func dogHandler(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	data := data{
		DataType: "Dog",
		Data:     name,
	}
	tpl.Execute(res, data)
}

func meHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("request object", req)
	url_query := req.URL.Query()
	name := url_query.Get("name")
	log.Println("Passed name param value", name)
	header_val := req.Header.Get("X-Custom-Header")
	log.Println("Header value ", header_val)
	data := data{
		DataType: "Me",
		Data:     name,
	}
	tpl.Execute(res, data)
}

type data struct {
	DataType string
	Data     string
}

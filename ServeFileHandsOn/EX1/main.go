/*
ListenAndServe on port 8080 of localhost:
For the default route "/" Have a func called "foo" which writes to the response "foo ran"
For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "
This is from dog
and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
*/

package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/pic", pic)
	http.HandleFunc("/pic.jpg", Servepic)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Foo ran")
}

func pic(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func Servepic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pic.jpg")
}

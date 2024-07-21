/*
Serve the files in the "starting-files" folder
To get your images to serve, use only this:
fs := http.FileServer(http.Dir("public"))
Hint: look to see what type FileServer returns, then think it through.
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("Templates/index.gohtml"))
}

func main() {
	/*
		Function Handle accepts Handler type as parameter,
		differ with HandleFunc which accepts “function” with 2 arguments which are ResponseWriter and http. Request.
	*/
	http.Handle("/pics/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8000", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Something went wrong during template execution")
	}
}

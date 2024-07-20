/*
Serve the files in the "starting-files" folder
Use "http.FileServer"
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("."))))
}

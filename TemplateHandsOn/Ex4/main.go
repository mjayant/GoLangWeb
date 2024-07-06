/*
Problem statement :
Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
Parse an html template, then pass the data from step 1 into the CSV template;
 have the html template display the CSV data in a web page.
*/

package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	http.HandleFunc("/", handlerFunction)
	http.ListenAndServe(":8000", nil)
}

type record struct {
	Date time.Time
	Open float64
}

func handlerFunction(res http.ResponseWriter, req *http.Request) {
	records := CSVParser("table.csv")
	tpl.Execute(res, records)

}

func CSVParser(filePath string) []record {
	var records []record
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	for _, row := range rows {
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		re := record{
			Date: date,
			Open: open,
		}
		records = append(records, re)
	}

	return records
}

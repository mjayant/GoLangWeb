/*
Problem statement:
Create a data structure to pass to a template which
contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
*/
package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {

	item1 := foodItem{
		Name:     "Breakfast Rice",
		ItemType: "Breakfast",
	}

	item2 := foodItem{
		Name:     "Lunch Rice",
		ItemType: "Lunch",
	}

	item3 := foodItem{
		Name:     "Dinner Rice",
		ItemType: "Dinner",
	}

	menu := Menu{
		Breakfast: []foodItem{item1},
		Lunch:     []foodItem{item2},
		Dinner:    []foodItem{item3},
	}
	err := tpl.Execute(os.Stdout, menu)
	if err != nil {
		log.Fatal("Something went wrong ", err)
	}
}

type foodItem struct {
	Name     string
	ItemType string
}

type Menu struct {
	Breakfast []foodItem
	Lunch     []foodItem
	Dinner    []foodItem
}

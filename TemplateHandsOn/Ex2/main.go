/*
Problem statemnet :
Create a data structure to pass to a template which
contains information about California hotels including Name, Address, City, Zip, Region
region can be: Southern, Central, Northern
can hold an unlimited number of hotels
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
	h1 := hotel{
		Name:    "Southern ABC Hotel",
		Address: "Southeren street 123",
		City:    "California",
	}

	h2 := hotel{
		Name:    "Southern DEF Hotel",
		Address: "Southeren street DEF",
		City:    "California",
	}

	h11 := hotel{
		Name:    "Central ABC Hotel",
		Address: "Central street ABC",
		City:    "California",
	}
	h22 := hotel{
		Name:    "Central DEF Hotel",
		Address: "Central street DEF",
		City:    "California",
	}

	h111 := hotel{
		Name:    "Northern ABC Hotel",
		Address: "Northern street ABC",
		City:    "California",
	}
	h222 := hotel{
		Name:    "Northern DEF Hotel",
		Address: "Northern street DEF",
		City:    "California",
	}
	region1 := region{
		Region: "Southern",
		Zip:    "1234",
		Hotels: []hotel{h1, h2},
	}

	region2 := region{
		Region: "Central",
		Zip:    "1234",
		Hotels: []hotel{h11, h22},
	}
	region3 := region{
		Region: "Northern",
		Zip:    "1234",
		Hotels: []hotel{h111, h222},
	}

	city_cal := city{
		Name:   "California",
		Region: []region{region1, region2, region3},
	}
	err := tpl.Execute(os.Stdout, city_cal)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
}

type hotel struct {
	Name    string
	Address string
	City    string
}

type region struct {
	Region string
	Zip    string
	Hotels []hotel
}

type city struct {
	Name   string
	Region []region
}

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
	c1 := course{Number: "A-123",
		Name:  "Math123",
		Units: "5"}

	c2 := course{Number: "B-123",
		Name:  "Scince123",
		Units: "5"}

	c3 := course{Number: "c-123",
		Name:  "Arts123",
		Units: "5"}

	s1 := semester{
		Term:    "Fall",
		Courses: []course{c1, c2, c3},
	}

	c11 := course{Number: "A-456",
		Name:  "Math456",
		Units: "5"}

	c22 := course{Number: "B-456",
		Name:  "Scince456",
		Units: "5"}

	c33 := course{Number: "c-456",
		Name:  "Arts456",
		Units: "5"}
	s11 := semester{
		Term:    "Spring",
		Courses: []course{c11, c22, c33},
	}

	c111 := course{Number: "A-789",
		Name:  "Math789",
		Units: "5"}

	c222 := course{Number: "B-789",
		Name:  "Scince789",
		Units: "5"}

	c333 := course{Number: "c-789",
		Name:  "Arts789",
		Units: "5"}

	s111 := semester{
		Term:    "Summer",
		Courses: []course{c111, c222, c333},
	}

	year1 := year{
		AcaYear: "2024",
		Fall:    s1,
		Spring:  s11,
		Summer:  s111,
	}
	years := []year{year1}
	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatal(err)
	}
}

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

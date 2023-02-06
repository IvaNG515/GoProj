package main

import "fmt"

type contactInfo struct {
	email    string
	zip_code int
}

type person struct {
	first_Name string
	last_Name  string
	contactInfo
}

func main() {
	jim := person{
		first_Name: "Jim",
		last_Name:  "Henderson",
		contactInfo: contactInfo{
			email:    "jimHenderson@email.com",
			zip_code: 1234567,
		},
	}
	jim.updateName("jimm")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFN string) {
	p.first_Name = newFN
}

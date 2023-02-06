package main

import "fmt"

type contactInfo struct {
	email    string
	zip_code int
}

type person struct {
	first_Name string
	last_Name  string
	contact    contactInfo
}

func main() {
	jim := person{
		first_Name: "Jim",
		last_Name:  "Henderson",
		contact: contactInfo{
			email:    "jimHenderson@email.com",
			zip_code: 1234567,
		},
	}
	fmt.Printf("somthing %+v", jim)
}

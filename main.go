package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	a := person{firstName: "alex", lastName: "brown"}
	fmt.Println(a)

	jim := person{
		firstName: "Jim",
		lastName:  "Carry",
		contact: contactInfo{
			email:   "jim.carry@gmail.com",
			zipCode: 94122,
		},
	}

	fmt.Printf("%+v", jim)

}

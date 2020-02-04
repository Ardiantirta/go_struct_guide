package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	a := person{firstName: "alex", lastName: "brown"}
	fmt.Println(a)

	jim := person{
		firstName: "Jim",
		lastName:  "Carry",
		contactInfo: contactInfo{
			email:   "jim.carry@gmail.com",
			zipCode: 94122,
		},
	}

	jim.updateName("gg")
	jim.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}

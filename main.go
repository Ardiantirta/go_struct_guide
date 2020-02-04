package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	a := person{firstName: "alex", lastName: "brown"}
	fmt.Println(a)

}

package main

import "fmt"

type contanctInfo struct {
	email string 
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contanctInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contanctInfo: contanctInfo{
			email: "jim@gmail.com",
			zipCode: 9400,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
	
}

func (p person) updateName (newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
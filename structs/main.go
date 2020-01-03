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
	// &variable equals to accessing the memory address of this variable
	// so then we can overwrite this and our code runs as desired
	// it is also pointer to a type person
	// jimPointer := &jim 
	// jimPointer.updateName("Jimmy")
	// jim.print()

	jim.updateName("Jimmy")
	jim.print()
	
}

// we can only use this function with a type pointerToPerson 
// which jim is actually
func (pointerToPerson *person) updateName (newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
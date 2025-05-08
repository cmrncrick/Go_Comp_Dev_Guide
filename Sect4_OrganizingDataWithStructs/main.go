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
	// // One way to assign person
	// // alex := person{"Alex", "Anderson"}

	// // Another way == Best Way
	// // alex := person{firstName: "Alex", lastName: "Anderson"}

	// // Another way
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)

	// // Will print out all keys with their values
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()

	// & is used to give memory address of the value the variable is pointing at
	// &jim gets turned into memory address (pointer)
	// this is not necessary, however.
	// jimPointer := &jim
	
	jim.updateName("Jimmy")
	jim.print()
}

// * is used to give the value the memory address is pointing at
// *person is a pointer that points to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// pointerTorPerson is the memory address Jim lives at
	// so the * operator is turning the memory address back into the value
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
package main

/*
Structs very important part about Go
Like classes since we have no classes
Create a strcut for a person
	- name
	- age
	- email

Different methods associated with the struct
	- Pointer receivers
	- Value receivers
*/
import (
	"fmt"
	"strconv"
)

// Person struct type defined
type Person struct {
	/* firstName string
	lastName  string
	city      string
	gender    string
	age       int */
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver) performs a caluclations and don't change anything
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver) we are going to change something
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)   we are going to change something
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Natlaia", lastName: "Mata", city: "Broomfield", gender: "F", age: 19}
	// Alternative
	person2 := Person{"Bob", "Johnson", "Pueblo", "M", 25}

	/* fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1) */

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}

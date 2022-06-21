package main

import "fmt"

const (
	first  = iota + 6
	second = 2 << iota // don't have to declare iota everytime, will reuse constant expresion from above line
)

const (
	third = iota // const blocks will reset back to zero
	fourth
)

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Variables and Data Types -*-*-*-*-*-*-*-*-*-*-*")
	var i int
	i = 42 // Long syntax
	fmt.Println(i)

	var f float32 = 3.14 // declaration with assignment
	fmt.Println(f)

	firstName := "Mario" // implicit inticialization and will  mostly see this in Go
	fmt.Println(firstName)

	boolTest := true
	fmt.Println(boolTest)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Pointers -*-*-*-*-*-*-*-*-*-*-*")
	var FirstName *string = new(string) // pointer operator

	*FirstName = "Mario" // de-reference operator

	fmt.Println(*FirstName) // nil represents a empty pointer

	LastName := "Zamora"
	fmt.Println(LastName)

	ptr := &LastName // address of operator with &&&&& symbol
	fmt.Println(ptr, *ptr)

	LastName = "Salais"
	fmt.Println(ptr, *ptr)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Constants -*-*-*-*-*-*-*-*-*-*-*")
	const pi = 3.1415 // have to declare at the same time we initialize
	fmt.Println(pi)

	const co int = 3
	fmt.Println(co + 3)
	// a bunch of random code
	fmt.Println(float32(co) + 1.2)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Iota and Constant Expressions -*-*-*-*-*-*-*-*-*-*-*")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)

}

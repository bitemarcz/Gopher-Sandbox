package main

import "fmt"

func main() {
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

}

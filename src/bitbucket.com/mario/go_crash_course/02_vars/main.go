package main

import (
	"fmt"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint unit8 unit16 unit32 unit64 uintptr
	// byte - alias for unit8
	// rune - alias for int32
	// float32 float64***
	// complex64 complex128

	// Using var
	var bradsname string = "Brad"
	var name = "Elizabeth"
	fmt.Println(bradsname)
	fmt.Println(name)
	var age int32 = 37  // Can be re-assigned or declared to another type or value
	const isCool = true // won't be able to re-assign the variable

	//shorthand - only works in function not global space
	// name := "Mario"
	size := 1.3

	name, email := "Mar", "mario.zamora@keysight.com"

	fmt.Println(name, age, email, isCool)
	fmt.Printf("%T\n", size)
}

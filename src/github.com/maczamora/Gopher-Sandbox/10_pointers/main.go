package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// Use pointers to pass alot of data stored at an address. If you choose to pass the address instead of data it can increase performance
	// Change values at specific locations, everthing passed by value
	// Not the best explanation
}

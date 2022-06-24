package main

import "fmt"

type HTTPRequest struct {
	Method string
}

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Loops on Condition  -*-*-*-*-*-*-*-*-*-*-*")
	var i int
	for i < 5 {
		println(i) // builtin function for debugging
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Conditional Loops with Post Clauses  -*-*-*-*-*-*-*-*-*-*-*")
	for ic := 0; ic < 5; ic++ { // implicit initialization variable only works within the for loop
		println(ic) // builtin function for debugging
	}

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Infinite Loops  -*-*-*-*-*-*-*-*-*-*-*")
	var ik int
	for { // implicit initialization variable only works within the for loop
		if ik == 5 {
			break
		}
		println(ik) // builtin function for debugging
		ik++
	}

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Loop over Collections  -*-*-*-*-*-*-*-*-*-*-*")
	slice := []int{1, 2, 3}
	for _, i := range slice {
		println(i)
	}

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Working with the Panic Function  -*-*-*-*-*-*-*-*-*-*-*")
	println("Starting web server")

	// do important things
	// panic("WTF Happened???")

	println("Web server started")

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* If Statements  -*-*-*-*-*-*-*-*-*-*-*")
	type User struct {
		ID        int
		firstName string
		lastName  string
	}

	u1 := User{
		ID:        1,
		firstName: "Wario",
		lastName:  "Zamora",
	}

	u2 := User{
		ID:        2,
		firstName: "Beth",
		lastName:  "Zamora",
	}

	if u1 == u2 {
		println("Same user!")
	} else if u1.firstName == u2.firstName {
		println("Similar Users")
	} else {
		println("Different Users ")
	}

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Switch Statements -*-*-*-*-*-*-*-*-*-*-*")
	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println(":Get request")
		fallthrough
	case "DELETE":
		println(":Delete request")
	case "POST":
		println(":POST request")
	case "PUT":
		println(":PUT request")
	default:
		println("Unhandled method")
	}
}

// First thing to be aware of is that every loop is going to be a for loop - 1 basic construction
// loop till condition
// loop till condition with post clause
// infinite loops
// loop over collections

package main

import "fmt"

func main() {
	// Defines a  map
	//emails := make(map[string]string)

	// Assign kv
	/* emails["Bob"] = "bob@gmail.com"
	emails["Amber"] = "amber@gmail.com"
	emails["Mike"] = "mike@gmail.com" */

	// Declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Amber": "amber@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	delete(emails, "Bob")
	fmt.Println(emails)
}

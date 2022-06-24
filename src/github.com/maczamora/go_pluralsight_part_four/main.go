package main

import (
	"fmt"
)

// arguments and parameters are basically the same
// parameters is on the function declaration side
// providing the data into the function is the argument

func main() {
	fmt.Println("Hello Playground!")
	port := 3000
	_, err := startWebserver(port, 2) // write only variable dump any type of data into this variable and don't need to use it
	fmt.Println(err)
}

func startWebserver(port, numberOfRetries int) (int, error) { // multiple returns in parans specify data you want returned
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	// return errors.New("Something went wrong") errors package that returns errors
	return port, nil
}

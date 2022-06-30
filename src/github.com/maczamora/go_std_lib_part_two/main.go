package main

import (
	"bufio"
	"fmt"
	"os"
)

type (
	messageType int
)

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Using the fmt package-*-*-*-*-*-*-*-*-*-*-*")
	// Sscanf is reverse of PrintF
	// fmt.Print("Waht is your name?")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Printf("Hello %s nice to meet you\n", name)

	// file, err := os.Open("text.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	var age int
	// 	var name string

	// 	n, err := fmt.Sscanf(scanner.Text(), "%s is %d years old", &name, &age)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if n == 2 {
	// 		fmt.Printf("%s,%d\n", name, age)
	// 	}
	// }
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Formatting Output -*-*-*-*-*-*-*-*-*-*-*")
	var age = 42
	var name = "Jeremey"
	var out, _ = fmt.Printf("My name is %s and I am %d years old\n", name, age)
	print("Bytes written out: ", out, "\n")

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Formatting Output -*-*-*-*-*-*-*-*-*-*-*")
	filename := "test.txt"
	showMessage(INFO, fmt.Sprintf("About to open %s", filename))
	showMessage(WARNING, fmt.Sprintf("If the file is not present, this application will fail "))

	file, err := os.Open(filename)
	if err != nil {
		print(err)
		showMessage(ERROR, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Formatting Other Data Types -*-*-*-*-*-*-*-*-*-*-*")

}

func showMessage(messagetype messageType, message string) {
	switch messagetype {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: \n%s\n", message)
		fmt.Printf(InfoColor, printMessage)
	}
}

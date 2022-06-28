package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* CLI Apps-*-*-*-*-*-*-*-*-*-*-*")
	// this is where stuff happens
	fmt.Println("Our current version of Go is " + runtime.Version()) // with carriage return
	fmt.Printf("We are using Go %v! running in %v\n", runtime.Version(), runtime.GOOS)

	// reader := bufio.NewReader(os.Stdin) // Grabbing values form the end-user
	// fmt.Println("What is your name?")
	// text, _ := reader.ReadString('\n')
	// fmt.Printf("Hello %v", text) // formats values in the output
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* CLI Arguments-*-*-*-*-*-*-*-*-*-*-*")
	// parsing aruments
	// args := os.Args[1:] // slice to take in a list of arguemtns
	// if len(args) == 1 && args[0] == "/help" {
	// 	fmt.Println("Usage: dinnertotal <Total Meal Amount> <Tip Percentage>")
	// 	fmt.Println("Example: dinnertotal 20 10")

	// } else {

	// 	if len(args) != 2 {
	// 		fmt.Println("You must enter 2 arguments! type /help for help")
	// 	} else {
	// 		mealTotal, _ := strconv.ParseFloat(args[0], 32)
	// 		tipAmount, _ := strconv.ParseFloat(args[1], 32)
	// 		fmt.Printf("\n Your meal total will be %.2f\n", calculateTotal(float32(mealTotal), float32(tipAmount)))
	// 	}
	// }

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Adding Flags To The App-*-*-*-*-*-*-*-*-*-*-*")
	// archPtr := flag.String("arch", "x86", "CPU Type")

	// flag.Parse()

	// switch *archPtr {
	// case "x86":
	// 	fmt.Println("Running in 32 bit mode")
	// case "AMD64":
	// 	fmt.Println("Running in 64 bit mode")
	// case "IA64":
	// 	fmt.Println("Remember IA64?")

	// }

	// fmt.Println("Process complete")

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Reading Keyboard Inputs-*-*-*-*-*-*-*-*-*-*-*")
	// var name string
	// fmt.Println("What is your name?")
	// inputs, _ := fmt.Scanf("%q", &name)

	// switch inputs {
	// case 0:
	// 	fmt.Printf("You must enter a name!\n")
	// case 1:
	// 	fmt.Printf("Hello %s! Nice to meet you\n", name)
	// }

	// fmt.Printf("Hello %s!! You input %d\n", name, inputs)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Intro to Bufio-*-*-*-*-*-*-*-*-*-*-*")
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // interactive probram or text based game or reading text in real time
		fmt.Println(scanner.Text())
		// do something while scanning
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// func calculateTotal(mealTotal float32, tipAmount float32) float32 {
// 	totalPrice := mealTotal + (mealTotal * (tipAmount / 100))
// 	return totalPrice
// }

// Why use the standard library??
// prewritten and tested code
// maintained by Go contributors
// backward Compatibility
// dependencies
// portability

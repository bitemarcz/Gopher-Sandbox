package main

import "fmt"

func main() {

	//var colors map[string]string // empty map - python dict

	// colors := make(map[int]string)

	//Example of making a map with key/value pairs
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#tg45670",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

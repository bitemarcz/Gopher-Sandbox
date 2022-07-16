package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Formatting Time Output -*-*-*-*-*-*-*-*-*-*-*")

	first := time.Now()

	// Year := t.Year()
	// Month := t.Month()
	// Day := t.Day()

	// fmt.Printf("Today is %d/%d/%d\n", Month, Day, Year)

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Calculating Time Spans -*-*-*-*-*-*-*-*-*-*-*")
	fmt.Printf("It is currently %v\n", first.Format("15:04:05"))
	time.Sleep(2 * time.Second)
	second := time.Now()
	fmt.Printf("It is now %v\n", second.Format("15:04:05"))

	fmt.Println("-*-*-*-*-*-*-*-*-*-*-* Calculating Elapsed Time for Applications -*-*-*-*-*-*-*-*-*-*-*")
	start := time.Now()
	args := os.Args
	// open the customer list
	custlist, err := os.Open(string(args[1]))
	check(err)
	defer custlist.Close()

	// create an output file
	outfile, err := os.Create("outfile.csv")
	check(err)
	defer outfile.Close()

	elapsed := time.Since(start)
	fmt.Println(elapsed)
	// scan the customer list
	scanner := bufio.NewScanner(custlist)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), ",")
		outfile.WriteString(names[1] + "," + names[2] + "\n")
	}
}

// Wall clock
// Used to keep track of the current time of day
// Subkect to variation
// Great for humans
// Not great for measureing time

// Monotonic Clock
// Used for measuring time - stop watches are a good example of this
// Not subject to variation
// Only meaningufl for the process calling it
// Can be simulated by wall clock time

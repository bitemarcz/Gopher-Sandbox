package main

import (
	"fmt"
	"math"

	"bitbucket.com/mario/go_crash_course/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleH"))
	fmt.Println(strutil.Reverse("mario alfredo zamora"))
}

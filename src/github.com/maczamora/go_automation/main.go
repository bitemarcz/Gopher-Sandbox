package main

import (
	"fmt"
	"math/rand"
)

func RandomString(n int) string {
	var random_characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = random_characters[rand.Intn(len(random_characters))]
	}
	return string(s)
}

func main() {
	fmt.Println(RandomString(30))
}

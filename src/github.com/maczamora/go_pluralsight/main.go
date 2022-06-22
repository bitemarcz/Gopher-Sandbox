package main

import (
	"fmt"

	"github.com/maczamora/go_pluralsight/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Liz",
		LastName:  "Beth",
	}
	fmt.Println(u)
}

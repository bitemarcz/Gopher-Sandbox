package main

import (
	"fmt"

	"github.com/maczamora/go_datatypes/organization"
)

func main() {
	p := organization.NewPerson("Mario", "Zamora")
	err := p.SetTwitterHandler("@marz")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	println(p.TwitterHandler())
	println(p.ID())
	println(p.FullName())
}

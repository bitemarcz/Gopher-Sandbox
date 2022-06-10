package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://uat.ironbridge.io/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Status code error:", err)
	} else {
		fmt.Println(resp)
	}

}

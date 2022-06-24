package main

import (
	"net/http"

	"github.com/maczamora/go_pluralsight/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

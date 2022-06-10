package main

import (
	"net/http"

	"bitbucket.com/mario/go_pluralsight/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

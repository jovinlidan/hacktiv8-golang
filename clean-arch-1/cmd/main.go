package main

import (
	"clean-arch-1/httpserver"
)

func main() {
	app := httpserver.CreateRouter()
	app.Run(":3001")
}

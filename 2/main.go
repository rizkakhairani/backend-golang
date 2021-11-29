package main

import (
	"backend-golang/2/routes"
)

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
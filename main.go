package main

import (
	"fmt"
	"pull-request-finder/app"
)

func main() {

	err := app.Curl()
	if err != nil {
		fmt.Print(err)
	}
}

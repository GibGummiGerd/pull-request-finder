package main

import (
	"fmt"
	"pr-finder/app"
)

func main() {

	err := app.Curl()
	if err != nil {
		fmt.Print(err)
	}
}

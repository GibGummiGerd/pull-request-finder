package main

import (
	"fmt"
	"os"
	"pull-request-finder/app"
)

func main() {

	err := app.PullRequestFlow()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Information successfully retrieved. Stored in pullRequests.yaml")
}

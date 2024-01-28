package main

import (
	"fmt"
	"net/url"
	"os"
)

/*
how to make a http get request
there's already a go server setup
going to write the client side, output http body
*/

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("ERR: missing argument \n Usage: go run main.go <url>")
		os.Exit(1)
	}

	// validate if the argument is a url
	_, err := url.ParseRequestURI(args[1])
	if err != nil {
		fmt.Printf("URL is not in a valid format %s\n: ", err)
		os.Exit(1)
	}

}

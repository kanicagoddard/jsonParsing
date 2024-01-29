// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// )

// /*
// how to make a http get request
// there's already a go server setup - dir/test-server
// going to write the client side, output http body
// */

// func main() {

// 	args := os.Args

// 	if len(args) < 2 {
// 		fmt.Printf("ERR: missing argument \n Usage: go run main.go <url>")
// 		os.Exit(1)
// 	}

// 	// validate if the argument is a url
// 	_, err := url.ParseRequestURI(args[1])
// 	if err != nil {
// 		fmt.Printf("URL is not in a valid format %s\n: ", err)
// 		os.Exit(1)
// 	}

// 	// get request
// 	response, err := http.Get(os.Args[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	// read body response
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("HTTP Status Code: %d\n Body: %v\n ", response.StatusCode, body)
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Post represents the structure of our JSON data
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// URL of the JSONPlaceholder API
	url := "https://jsonplaceholder.typicode.com/posts"

	// Fetch the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching data:", err)
	}
	defer resp.Body.Close()

	// Read and parse the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body:", err)
	}

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	// Display the posts
	for _, post := range posts {
		words := strings.Fields(post.Body)
		if len(words) > 5 {
			words = words[:5]
			// fmt.Printf("ID: %d, Title: %s, UserID: %d \n", post.ID, post.Body, post.UserID)
			fmt.Printf("ID: %d, Body: %s...\n", post.ID, strings.Join(words, " "))
		}
	}
}

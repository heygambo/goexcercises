// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/"
// "/dog/"
// "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "homepage")
	// alternative
	// io.WriteString(w, "homepage")
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "dog")
	// alternative
	// io.WriteString(w, "dog")
}

func handleMe(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "It's a me")
	// alternative
	// io.WriteString(w, "It's a me")
}

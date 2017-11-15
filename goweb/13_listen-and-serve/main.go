// # ListenAndServe on port 8080 of localhost

// For the default route "/"
// Have a func called "foo"
// which writes to the response "foo ran"

// For the route "/dog/"
// Have a func called "dog"
// which parses a template called "dog.gohtml"
// and writes to the response "<h1>This is from dog</h1>"
// and also shows a picture of a dog when the template is executed.

// Use "http.ServeFile"
// to serve the file "dog.jpeg"

package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/toby.jpg", toby)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "I ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln("Error loading template", err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func toby(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}

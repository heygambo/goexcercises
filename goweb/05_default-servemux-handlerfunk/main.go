// 1. Take the previous program and change it so that:
// * func main uses http.Handle instead of http.HandleFunc

// Contstraint: Do not change anything outside of func main

// Hints:

// [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
// ``` Go
// type HandlerFunc func(ResponseWriter, *Request)
// ```

// [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
// ``` Go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// ```

// [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
// ``` Go
//   func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
//   		mux.Handle(pattern, HandlerFunc(handler))
//   }
// ```

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("homepage.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(handleHome))
	http.Handle("/dog/", http.HandlerFunc(handleDog))
	http.Handle("/me/", http.HandlerFunc(handleMe))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	data := struct { Title string }{ "Homepage" }
	fmt.Println("data", data.Title)
	err := tpl.ExecuteTemplate(w, "homepage.gohtml", data)
	if err != nil {
		log.Fatalln("error while executing template", err)
	}
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

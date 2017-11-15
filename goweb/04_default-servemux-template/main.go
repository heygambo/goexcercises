// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template

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
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}

func handleHome(w http.ResponseWriter, req *http.Request) {
	data := struct{ Title string }{"Homepage"}
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

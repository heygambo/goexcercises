// # Serve the files in the "starting-files" folder

// ## To get your images to serve, use only this:

// ``` Go
// 	fs := http.FileServer(http.Dir("public"))
// ```

// Hint: look to see what type FileServer returns, then think it through.

package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./starting-files/templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("./starting-files/public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Error while executing template:", err)
	}
}

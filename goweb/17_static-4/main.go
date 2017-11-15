// # Serve the files in the "starting-files" folder

// ## To get your images to serve, use:

// ``` Go
// 	func StripPrefix(prefix string, h Handler) Handler
// 	func FileServer(root FileSystem) Handler
// ```

// Constraint: you are not allowed to change the route being used for images in the template file

package main

import (
	"log"
	"net/http"
)

func main() {
	// fs := http.StripPrefix("./starting-files", http.FileServer(http.Dir("./starting-files")))
	// ?? not needed
	fs := http.FileServer(http.Dir("./starting-files"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

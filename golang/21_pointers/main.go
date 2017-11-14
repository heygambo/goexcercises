package main

import (
	"fmt"
	"text/template"
)

var tpl1 *template.Template

func main() {
	fmt.Printf("%p", tpl1)
}

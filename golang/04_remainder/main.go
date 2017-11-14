package main

import (
	"fmt"
)

func main() {
	var small, large int

	fmt.Println("I'll print a reminder.")

	fmt.Print("Enter a small number: ")
	fmt.Scan(&small)

	fmt.Println("")

	fmt.Print("Enter a large number: ")
	fmt.Scan(&large)

	fmt.Println("")

	remainder := large % small
	fmt.Println("The remainder is:", remainder)
}

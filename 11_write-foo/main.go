package main

import (
	"fmt"
)

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(numbers ...int) {
	if len(numbers) <= 0 {
		fmt.Println("No numbers given.")
		return
	}
	fmt.Println("Printing:", numbers)
	for _, v := range numbers {
		fmt.Println("number:", v)
	}
}

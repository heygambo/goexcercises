package main

import (
	"fmt"
)

func main() {
	fmt.Println("Highest Number:", highestNumber(3, 400, 5, 6))
}

func highestNumber(numbers ...float64) float64 {
	var highest float64
	for _, n := range numbers {
		if n > highest {
			highest = n
		}
	}
	return highest
}

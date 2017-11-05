// This program aims to solve this problem:
// https://projecteuler.net/problem=5

package main

import (
	"fmt"
)

const maxCount = 1000000000
const highestDevisor = 20

func main() {
	for i := 1; i < maxCount; i++ {
		// fmt.Println("Testing number", i)
		if isEvenlyDivisable(i) {
			fmt.Println("Found the number:", i)
			return
		}
	}
}

func isEvenlyDivisable(n int) bool {
	for i := 1; i <= highestDevisor; i++ {
		if n%i > 0 {
			return false
		}
	}
	return true
}

/*
This program aims to solve this problem:
https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"
)

const maxCount = 1000000000
const highestDevisor = 25

func main() {
	number := make(chan int)
	done := make(chan bool)


	
	for i := 1; i < maxCount; i++ {
		// fmt.Println("Testing number", i)
		go func () {
			if isEvenlyDivisable(i) {
				fmt.Println("Found the number:", i)
				return
			}
		}()
	}

	<-done
}

func isEvenlyDivisable(n int) bool {
	for i := 1; i <= highestDevisor; i++ {
		if n%i > 0 {
			return false
		}
	}
	return true
}

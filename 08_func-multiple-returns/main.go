package main

import (
	"fmt"
)

func main() {
	n, b := divider(32)
	fmt.Println("n / 2 =", n)
	fmt.Println("isEven", b)
}

func divider(n int) (float64, bool) {
	divided := float64(n) / 2.0
	isEven := n%2 == 0
	return divided, isEven
}

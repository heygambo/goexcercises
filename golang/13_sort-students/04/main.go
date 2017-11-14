package main

import "sort"
import "fmt"

func main() {
	n := []int{6, 1, 67, 87, 3, 45325, 234, 63625}
	sort.Ints(n)
	fmt.Println(n)
}

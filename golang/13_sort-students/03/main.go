package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(people)
	fmt.Println(people)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"Zeno", "John", "Al", "Jenny"}
	sort.StringSlice(studyGroup).Sort()
	fmt.Println(studyGroup)
}

package main

import (
	"fmt"
	"strconv"
)

var count int

func main() {
	c := incrementor(2)
	for s := range c {
		count++
		fmt.Println(s)
	}
	fmt.Println("Final Counter:", count)
}

func incrementor(n int) <-chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprintf("Process: " + strconv.Itoa(i) + " printing: " + strconv.Itoa(k))
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/

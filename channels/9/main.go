package main

import (
	"fmt"
)

func main() {
	c := generate(4, 10)

	// fan out
	d1 := divide(c)
	d2 := divide(c)

	// fan in

}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}
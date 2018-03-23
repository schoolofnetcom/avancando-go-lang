package main

import (
	"fmt"
)

// Pipeline
func main() {
	numbers := generate(2, 4, 6)
	result := divide(numbers)

	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()
	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)
	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()
	return channel
}
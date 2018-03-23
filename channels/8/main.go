package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {

	x := funnel(generateMsg("hello go"), generateMsg("hello world"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}
	fmt.Println("Finished...")

}

func generateMsg(s string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("String: %s - Value: %d", s, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()
	return channel
}

func funnel(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- <-channel1
		}
	}()

	go func() {
		for {
			channel <- <-channel2
		}
	}()
	return channel
}
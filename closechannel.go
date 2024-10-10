package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int, 2)

	go func() {
		time.Sleep(5 * time.Second)
		channel <- 1
		channel <- 2
		close(channel)
	}()

	fmt.Println("before")
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("after")
}

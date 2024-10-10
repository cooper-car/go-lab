package main

import (
	"fmt"
	"sync"
	"time"
)

// 知道次數
func main() {

	channel := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			channel <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		wg.Done()
	}()

	wg.Wait()
}

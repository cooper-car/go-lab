package main

import (
	"fmt"
	"sync"
	"time"
)

// for range 搭配 close channel 來使用
func main() {

	channel := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			channel <- i
		}
		close(channel)
	}()

	wg.Add(1)
	go func() {
		for i := range channel {
			fmt.Println(i)
		}
		wg.Done()
	}()

	wg.Wait()
}

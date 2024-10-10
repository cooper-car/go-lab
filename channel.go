package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("before sending data")
		time.Sleep(5 * time.Second)
		ch <- 1
		fmt.Println("after sending data")
	}()

	// 主進程會等工作進程完成後才會繼續執行
	fmt.Println("waiting for data")
	num := <-ch
	fmt.Println("received data: ", num)
}

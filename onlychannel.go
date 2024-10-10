package main

import (
	"fmt"
	"sync"
)

// v1
//func main() {
//	channel := make(chan int)
//	var wg sync.WaitGroup
//
//	wg.Add(1)
//	go func() {
//		channel <- 1
//		wg.Done()
//	}()
//
//	fmt.Print(<-channel)
//	wg.Wait()
//}

// v2
func main() {
	channel := make(chan int)
	var wg sync.WaitGroup
	var wgRead sync.WaitGroup

	wg.Add(1)
	go func() {
		channel <- 1
		wg.Done()
	}()

	wgRead.Add(1)
	go func() {
		for v := range channel {
			fmt.Print(v)
		}
		wgRead.Done()
	}()

	wg.Wait()
	close(channel)
	wgRead.Wait()
}

/*
如果 fmt.Print(<-channel) 寫在 wg.Wait() 之後，可能會導致 go func() 已經結束並且通道已經被關閉，從而無法從通道中讀取值。
因此，必須在等待 wg.Wait() 之前讀取通道。

如果用 for range 读取通道，記得加上 close(channel) 來關閉通道，否則會造成 deadlock。
補上 wgRead.Add(1) 來計算 goroutine 數目，並加上 wgRead.Wait() 來等待 goroutine 完成。
*/

package main

import (
	"log"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var lock sync.RWMutex

	wg.Add(2)
	go func() {
		log.Println("A: Lock before")
		lock.RLock()

		log.Println("A: Locked")
		time.Sleep(5 * time.Second)

		log.Println("A: Unlocked")
		lock.RUnlock()

		wg.Done()
	}()

	go func() {
		log.Println("B: Lock before")
		lock.RLock()

		log.Println("B: Locked")
		time.Sleep(5 * time.Second)

		log.Println("B: Unlocked")
		lock.RUnlock()

		wg.Done()
	}()

	wg.Wait()
}

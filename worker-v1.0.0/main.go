package main

import (
	"fmt"
	"pool-test.com/pool"
	"sync"
)

func main() {

	worker := pool.NewPool(2)

	worker.AddTask(func() interface{} {
		return 1
	})

	worker.AddTask(func() interface{} {
		return 2
	})

	worker.AddTask(func() interface{} {
		return 3
	})

	worker.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for result := range worker.Results {
			fmt.Println(result)
		}
		wg.Done()
	}()
	worker.Wait()
	wg.Wait()
}

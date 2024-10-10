package main

import (
	"fmt"
	"sync"
	"worker-v2.com/pool"
)

func main() {

	worker := pool.NewPool(2)

	// 寫一個創建 Task 的函數
	createTask := func(i int) pool.Task {
		return func() interface{} {
			return i
		}
	}

	// 添加 10 個任務
	for i := 0; i < 10; i++ {
		worker.AddTask(createTask(i))
	}

	worker.Run()

	// 打印結果
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for result := range worker.Result {
			fmt.Println(result)
		}
		wg.Done()
	}()

	worker.Wait()
	wg.Wait()

}

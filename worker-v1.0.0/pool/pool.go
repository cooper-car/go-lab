package pool

import (
	"math"
	"sync"
)

type Task func() interface{}

type Pool struct {
	worker  int
	tasks   *Queue
	event   chan struct{}
	Results chan interface{}
	wg      sync.WaitGroup
}

func NewPool(worker int) *Pool {
	return &Pool{
		worker:  worker,
		tasks:   NewQueue(-1),
		event:   make(chan struct{}, math.MaxInt32),
		Results: make(chan interface{}, worker*3),
	}
}

// 向任務池添加任務
func (p *Pool) AddTask(task Task) {
	p.tasks.Append(task)
	p.event <- struct{}{}
}

func (p *Pool) Start() {
	for i := 0; i < p.worker; i++ {
		p.wg.Add(1)

		go func() {
			// 忽略從管道中讀取的數據
			for range p.event {

				// 從任務隊列中讀取任務
				task, err := p.tasks.Front()
				if err != nil {
					continue
				}

				// 將佇列中空介面資料轉換為Task並進行執行
				if task, ok := task.(Task); ok {
					// 將結果放入到 result 管道
					p.Results <- task()
				}
			}
			p.wg.Done()
		}()
	}
}

func (p *Pool) Wait() {
	close(p.event)
	p.wg.Wait()
	close(p.Results)
}

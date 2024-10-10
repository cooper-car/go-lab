package pool

import (
	"fmt"
	"sync"
)

const defaultCap = 10

type Queue struct {
	elements []interface{}
	locker   sync.Mutex
	limit    int
}

func NewQueue(limit int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, defaultCap),
		limit:    limit,
	}
}

func (q *Queue) Append(e interface{}) error {
	q.locker.Lock()
	defer q.locker.Unlock()

	if q.limit != -1 && len(q.elements) >= q.limit {
		return fmt.Errorf("queue is limit %d", q.limit)
	}
	q.elements = append(q.elements, e)

	return nil
}

func (q *Queue) Front() (interface{}, error) {
	if len(q.elements) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	// 移除第一個元素
	e := q.elements[0]
	q.elements = q.elements[1:]

	return e, nil
}

func (q *Queue) Len() int {
	return len(q.elements)
}

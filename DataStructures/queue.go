package DataStructures

import (
	"sync"
)

type Queue struct {
	items []any
	mu    sync.Mutex
}

func (q *Queue) Enqueue(item any) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() any {
	q.mu.Lock()
	defer q.mu.Unlock()
	last := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return last
}

func (q *Queue) Peek() any {
	return q.items[len(q.items)-1]
}

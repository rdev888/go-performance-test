package DataStructures

import "sync"

type Queue struct {
	items []any
	mu    sync.Mutex
}

func (q *Queue) enquue(item any) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
}

package DataStructures

import (
	"sync"
)

type Stack struct {
	items []any
	mx    sync.Mutex
}

func (s *Stack) Push(item any) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.items = append(s.items, item)
}

func (s *Stack) Pop() any {
	s.mx.Lock()
	defer s.mx.Unlock()
	if len(s.items) == 0 {
		return nil
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

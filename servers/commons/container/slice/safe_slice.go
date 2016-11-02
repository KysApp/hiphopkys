package slice

import (
	"sync"
)

type SafeSlice struct {
	sync.RWMutex
	Slice []interface{}
}

func New(length, capacity int) *SafeSlice {
	s := &SafeSlice{
		Slice: make([]interface{}, length, capacity),
	}
	return s
}

func (s *SafeSlice) Cap() int {
	s.RLock()
	c := cap(s.Slice)
	s.RUnlock()
	return c
}

func (s *SafeSlice) Len() int {
	s.RLock()
	l := len(s.Slice)
	s.RUnlock()
	return l
}

func (s *SafeSlice) Append(e interface{}) {
	s.Lock()
	s.Slice = append(s.Slice, e)
	s.Unlock()
}

func (s *SafeSlice) AppendArray(array ...interface{}) {
	s.Lock()
	s.Slice = append(s.Slice, array...)
	s.Unlock()
}

func (s *SafeSlice) Get(index int) interface{} {
	s.RLock()
	v := s.Slice[index]
	s.RUnlock()
	return v
}

func (s *SafeSlice) Remove(index int) {
	s.Lock()
	s.Slice = append(s.Slice[:index], s.Slice[index+1:]...)
	s.Unlock()
}

/**
 * [from,to)
 */
func (s *SafeSlice) Sub(from, to int) *SafeSlice {
	s.Lock()
	subSlice := s.Slice[from:to]
	s.Unlock()
	newSafeSlice := New(0, 0)
	newSafeSlice.Slice = subSlice
	return newSafeSlice
}

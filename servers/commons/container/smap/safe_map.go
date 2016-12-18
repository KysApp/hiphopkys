package smap

import (
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	Map map[interface{}]interface{}
}

func New() *SafeMap {
	m := &SafeMap{
		Map: make(map[interface{}]interface{}),
	}
	return m
}

func (m *SafeMap) Insert(k interface{}, v interface{}) {
	m.Lock()
	m.Map[k] = v
	m.Unlock()
}

func (m *SafeMap) Len() int {
	m.RLock()
	length := len(m.Map)
	m.RUnlock()
	return length
}

func (m *SafeMap) Remove(k interface{}) {
	m.Lock()
	delete(m.Map, k)
	m.Unlock()
}

func (m *SafeMap) Get(k interface{}) interface{} {
	m.RLock()
	v := (m.Map)[k]
	m.RUnlock()
	return v
}

func (m *SafeMap) IsExistence(k interface{}) bool {
	m.RLock()
	_, ok := (m.Map)[k]
	m.RUnlock()
	return ok
}

package smap

import (
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	Map map[string]interface{}
}

func New() *SafeMap {
	m := &SafeMap{
		Map: make(map[string]interface{}),
	}
	return m
}

func (m *SafeMap) Insert(k string, v interface{}) {
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

func (m *SafeMap) Remove(k string) {
	m.Lock()
	delete(m.Map, k)
	m.Unlock()
}

func (m *SafeMap) Get(k string) interface{} {
	m.RLock()
	v := (m.Map)[k]
	m.RUnlock()
	return v
}

func (m *SafeMap) IsExistence(k string) bool {
	m.RLock()
	_, ok := (m.Map)[k]
	m.RUnlock()
	return ok
}

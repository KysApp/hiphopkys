package list

import (
	"container/list"
	"sync"
)

type SafeList struct {
	sync.RWMutex
	List *list.List
}

func New() *SafeList {
	sl := &SafeList{
		List: list.New(),
	}
	return sl
}

func (l *SafeList) Back() *list.Element {
	l.RLock()
	e := l.List.Back()
	l.RUnlock()
	return e
}

func (l *SafeList) Front() *list.Element {
	l.RLock()
	e := l.List.Front()
	l.RUnlock()
	return e

}
func (l *SafeList) InsertAfter(v interface{}, mark *list.Element) *list.Element {
	l.Lock()
	e := l.List.InsertAfter(v, mark)
	l.Unlock()
	return e
}

func (l *SafeList) InsertBefore(v interface{}, mark *list.Element) *list.Element {
	l.Lock()
	e := l.List.InsertBefore(v, mark)
	l.Unlock()
	return e
}

func (l *SafeList) Len() int {
	l.RLock()
	length := l.List.Len()
	l.RUnlock()
	return length
}

func (l *SafeList) PushBack(v interface{}) *list.Element {
	l.Lock()
	e := l.List.PushBack(v)
	l.Unlock()
	return e
}

func (l *SafeList) PushBackList(other *SafeList) {
	l.Lock()
	other.RLock()
	l.List.PushBack(other.List)
	other.RUnlock()
	l.Unlock()
}

func (l *SafeList) PushFront(v interface{}) *list.Element {
	l.Lock()
	e := l.List.PushFront(v)
	l.Unlock()
	return e
}

func (l *SafeList) PushFrontList(other *SafeList) {
	l.Lock()
	other.RLock()
	l.List.PushFrontList(other.List)
	other.RUnlock()
	l.Unlock()
}

func (l *SafeList) Remove(e *list.Element) interface{} {
	l.Lock()
	re := l.List.Remove(e)
	l.Unlock()
	return re
}

func (l *SafeList) MoveAfter(e, mark *list.Element) {
	l.Lock()
	l.List.MoveAfter(e, mark)
	l.Unlock()
}

func (l *SafeList) MoveBefore(e, mark *list.Element) {
	l.Lock()
	l.List.MoveBefore(e, mark)
	l.Unlock()
}

func (l *SafeList) MoveToBack(e *list.Element) {
	l.Lock()
	l.List.MoveToBack(e)
	l.Unlock()
}

func (l *SafeList) MoveToFront(e *list.Element) {
	l.Lock()
	l.List.MoveToFront(e)
	l.Unlock()
}

func (l *SafeList) RemoveFirstElementWithValue(value interface{}) {
	var item *list.Element
	l.Lock()
	for v := l.List.Front(); v != nil; v = v.Next() {
		if value == v.Value {
			item = v
			break
		}
	}
	if item != nil {
		l.List.Remove(item)
	}
	l.Unlock()
}

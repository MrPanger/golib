package ds

import "sync"

type Set interface {
	Add(interface{}) bool
	Remove(interface{})
	IsExist(interface{}) bool
	Size() int
}

type set struct {
	lock  sync.RWMutex
	store map[interface{}]struct{}
}

func NewSet(cap int) Set {
	if cap < 0 {
		panic("cap can't be less than 0")
	}
	r := new(set)
	r.store = make(map[interface{}]struct{}, cap)
	return r
}

func (s *set) Add(i interface{}) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, exist := s.store[i]; exist {
		return false
	}
	s.store[i] = struct{}{}
	return true
}

func (s *set) Remove(i interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, exist := s.store[i]; exist {
		delete(s.store, i)
	}
}

func (s *set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.store)
}

func (s *set) IsExist(i interface{}) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, r := s.store[i]
	return r
}

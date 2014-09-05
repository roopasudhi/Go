package kv

import (
	"sync"
	"strings"
)
type Store interface{
	Get(k string)(string,bool)
	Put(k,v string)
	Delete(k string)
	Count(k string)int
}

type store struct{
	lock sync.RWMutex
	kv map[string]string
}

func (s *store) Get(k string) (string,bool){
	s.lock.RLock()
	defer s.lock.RUnlock()
	v,retrieved := s.kv[k]
	return v,retrieved
}

func (s *store) Put(k , v string){
	s.lock.Lock()
	defer s.lock.Unlock()
	s.kv[k] = v
}

func (s *store) Delete(k string){
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.kv,k)
} 

func (s *store) Count(f string) int{
	s.lock.RLock()
	defer s.lock.RUnlock()
	var cnt int
	for k := range s.kv{
		if strings.HasPrefix(k,f){
			cnt++
		}
	}
	return cnt
} 


func NewStore() *store{
	return &store{kv:make(map[string]string)}
}




package syncmap

import "sync"

type SyncMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

func NewSyncMap[K comparable, V any]() SyncMap[K, V] {
	return SyncMap[K, V]{m: make(map[K]V), mu: sync.RWMutex{}}
}

func (s *SyncMap[K, V]) Set(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[key] = value
}

func (s *SyncMap[K, V]) Get(key K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.m[key]
	return v, ok
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.m, key)
}

func (s *SyncMap[K, V]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.m)
}

func (s *SyncMap[K, V]) Exists(key K) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, ok := s.m[key]
	return ok
}

func (s *SyncMap[K, V]) Range(f func(K, V)) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for k, v := range s.m {
		f(k, v)
	}
}

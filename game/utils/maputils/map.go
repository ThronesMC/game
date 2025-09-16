package maputils

//Taken from: https://github.com/AkmalFairuz/df-game/blob/master/internal/sync_map.go

import (
	"iter"
	"sync"
)

type Map[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m: make(map[K]V),
	}
}

func (sm *Map[K, V]) Load(key K) (value V, ok bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok = sm.m[key]
	return
}

func (sm *Map[K, V]) MustLoad(key K) V {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.m[key]
}

func (sm *Map[K, V]) Store(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *Map[K, V]) LoadOrStore(key K, value V) (actual V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	actual, ok := sm.m[key]
	if !ok {
		sm.m[key] = value
		actual = value
	}
	return actual
}

func (sm *Map[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

func (sm *Map[K, V]) Len() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return len(sm.m)
}

func (sm *Map[K, V]) Map() iter.Seq2[K, V] {
	sm.mu.Lock()
	m := make(map[K]V, len(sm.m))
	for k, v := range sm.m {
		m[k] = v
	}
	sm.mu.Unlock()
	return func(yield func(K, V) bool) {
		for k, v := range m {
			if !yield(k, v) {
				break
			}
		}
	}
}

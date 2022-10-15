package mutex

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

func (s SafeMap[K, V]) LoadOrStore(key K, newVal V) (V, bool) {
	// 第一次判断
	oldVal, ok := s.get(key)
	if ok {
		return oldVal, ok
	}
	// 如果没有，加写锁
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// double check，二次校验
	oldVal, ok = s.m[key]
	if ok {
		return oldVal, true
	}
	s.m[key] = newVal
	return newVal, false
}

func (s SafeMap[K, V]) get(key K) (V, bool) {
	// 加读锁
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	oldVal, ok := s.m[key]
	return oldVal, ok
}

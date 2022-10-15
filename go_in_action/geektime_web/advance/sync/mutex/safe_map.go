package mutex

import "sync"

type SafeMap[K comparable, V any] struct {
	values map[K]V
	mutex  sync.RWMutex
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
	oldVal, ok = s.values[key]
	if ok {
		return oldVal, true
	}
	s.values[key] = newVal
	return newVal, false
}

func (s SafeMap[K, V]) LoadOrStoreV1(key K, newVal V) (V, bool) {
	s.mutex.RLock()
	oldVal, ok := s.values[key]
	/*
		这里如果用defer，那么到后面加写锁时这里的读锁还未释放，会造成死锁
		fatal error: all goroutines are asleep - deadlock!
	*/
	defer s.mutex.RUnlock()
	if ok {
		return oldVal, true
	}

	// 如果没有，加写锁
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// double check，二次校验
	oldVal, ok = s.values[key]
	if ok {
		return oldVal, true
	}
	s.values[key] = newVal
	return newVal, false
}

func (s SafeMap[K, V]) get(key K) (V, bool) {
	// 加读锁
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	oldVal, ok := s.values[key]
	return oldVal, ok
}

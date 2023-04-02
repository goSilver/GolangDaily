package a_singleton

import "sync"

type Singleton struct{}

var ins *Singleton

// GetInstanceLazy 懒汉式。非并发安全。
func GetInstanceLazy() *Singleton {
	if ins == nil {
		ins = &Singleton{}
	}
	return ins
}

// GetInstanceLazyWithLock 懒汉式加锁，并发安全，但是性能差
var mu sync.Mutex

func GetInstanceLazyWithLock() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if ins == nil {
		ins = &Singleton{}
	}
	return ins
}

// GetInstanceLazyWithDoubleCheck 懒汉式-双重检查。并发安全，性能ok。
// Java版该实现存在CPU指令重排序问题，需要volatile关键字禁止指令重排序
func GetInstanceLazyWithDoubleCheck() *Singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()

		if ins == nil {
			ins = &Singleton{}
		}
	}
	return ins
}

// GetInstanceHungry 饿汉式
var insHungry *Singleton = &Singleton{}

func GetInstanceHungry() *Singleton {
	return insHungry
}

func init() {
	if insHungry == nil {
		insHungry = new(Singleton)
	}
}

// GetInstanceOnce 利用go自带并发工具sync.Once
var once sync.Once

func GetInstanceOnce() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
	})
	return ins
}

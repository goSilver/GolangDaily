package main

import "sync"

func main() {

}

var mutex sync.Mutex
var rwMutex sync.RWMutex

func Mutex() {
	mutex.Lock()
	defer mutex.Unlock()

	// 业务代码
}

func RwMutex() {
	// 加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 加写锁
	rwMutex.Lock()
	defer mutex.Unlock()
}

func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

func Failed2() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

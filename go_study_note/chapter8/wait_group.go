package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 累加计数(原子操作)

		go func(id int) {
			defer wg.Done() // 递减计数
			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}

	println("main...")
	wg.Wait() // 阻塞，直到计数归零
	println("main exit...")
}

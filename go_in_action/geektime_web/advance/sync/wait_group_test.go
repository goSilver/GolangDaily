package sync

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"testing"
)

// TestWaitGroup waitGroup实践，1到10累加
func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	var result int64 = 0
	for i := 0; i < 10; i++ {
		// 计数加一
		wg.Add(1)
		// 启goroutine执行加法
		go func(delta int) {
			// 原子相加delta
			atomic.AddInt64(&result, int64(delta))
			// 执行相加后计数减一
			wg.Done()
		}(i)
	}
	// 阻塞等待所有goroutine执行完成
	wg.Wait()
	fmt.Println(result)
}

// TestErrGroup errgroup和waitGroup有同样到功能
func TestErrGroup(t *testing.T) {
	eg := errgroup.Group{}
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := i
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

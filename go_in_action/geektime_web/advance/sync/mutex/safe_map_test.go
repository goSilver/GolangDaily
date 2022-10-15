package mutex

import (
	"fmt"
	"testing"
	"time"
)

// TestDeferLock 验证单goroutine时不规范操作造成死锁
func TestDeferLock(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}
	sm.LoadOrStoreV1("a", "b")
}

func TestOverride(t *testing.T) {
	sm := SafeMap[string, string]{
		values: make(map[string]string, 4),
	}

	go func() {
		time.Sleep(time.Second)
		sm.LoadOrStore("a", "b")
	}()

	go func() {
		time.Sleep(time.Second)
		sm.LoadOrStore("a", "c")
	}()

	time.Sleep(time.Second)
	fmt.Println("Hello")
}

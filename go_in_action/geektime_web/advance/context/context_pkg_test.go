package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestContextWithValue 验证每一次调用WithValue方法返回的都是一个新的context实例
func TestContextWithValue(t *testing.T) {
	ctx := context.Background()
	// 设置键值对并返回一个新context parent
	parent := context.WithValue(ctx, "my_key", "my_value")
	// 设置键值对并返回一个新context children
	children := context.WithValue(ctx, "my_key", "my_new_vale")

	// 验证parent和children中的my_key是否一样
	fmt.Printf("parentCtx:%v \n", parent.Value("my_key"))
	fmt.Printf("childrenCtx:%v \n", children.Value("my_key"))
	/*
		打印：
		parentCtx:my_value
		childrenCtx:my_new_vale
	*/
}

// TestContextTimeout 父context超时后，子context也跟着结束，err字段存储了context结束原因
func TestContextTimeout(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancelFunc1 := context.WithTimeout(ctx, time.Second)
	subCtx, cancelFunc2 := context.WithTimeout(timeoutCtx, 3*time.Second)
	go func() {
		// 一秒钟之后就会过期，然后输出 timeout
		<-subCtx.Done()
		fmt.Println("timeout")
	}()
	time.Sleep(2 * time.Second)
	// 在这里打断点，debug到这里时context中到err字段是"context deadline exceeded"
	cancelFunc2()
	cancelFunc1()
}

// TestContextTimeoutExample
func TestContextTimeoutExample(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout")
	case <-bsChan:
		fmt.Println("business end")
	}
	/*
		打印：
		timeout
	*/
}

func slowBusiness() {
	time.Sleep(2 * time.Second)
}

// TestContextAfterFunc 实践AfterFunc
func TestContextAfterFunc(t *testing.T) {
	bsChan := make(chan struct{})
	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()

	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("timeout")
	})
	<-bsChan
	fmt.Println("business end")
	timer.Stop()
	/*
		打印：
		timeout
		business end
	*/
}

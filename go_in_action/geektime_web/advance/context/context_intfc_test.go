package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 实践context接口的方法

// TestContextValue 每一次调用WithValue方法返回的都是一个新的context实例
func TestContextValue(t *testing.T) {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "key", 123)
	val := childCtx.Value("key")
	fmt.Println(val)
	/*
		输出：
		123
	*/
}

// 父context cancel掉，子context也会结束
func TestParentCtx(t *testing.T) {
	ctx := context.Background()
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Minute))
	childCtx := context.WithValue(dlCtx, "key", 123)
	// 父context cancel掉
	cancel()
	err := childCtx.Err()
	fmt.Println(err)
	/*
		输出：
		context canceled
	*/
}

// TestParentValueCtx 子context拿不到父context的value，但是通过引用类型map可以
func TestParentValueCtx(t *testing.T) {
	ctx := context.Background()
	parentCtx := context.WithValue(ctx, "map", map[string]string{})
	childCtx := context.WithValue(parentCtx, "key1", "value1")
	// 类型断言
	m := childCtx.Value("map").(map[string]string)
	m["key1"] = "val1"

	// 父拿不到子的普通kv
	val := parentCtx.Value("key1")
	fmt.Println(val)
	// 引用类型可以拿到
	val = parentCtx.Value("map")
	fmt.Println(val)
	/*
		输出：
		<nil>
		map[key1:val1]
	*/
}

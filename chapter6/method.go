package main

import (
	"fmt"
)

type N int

// toString 普通方法
func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}

// test 方法内部不引用实例，可省略参数名
func (N) test() {
	println("hi!")
}

//
func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	(*n)++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func main() {
	var a N = 23
	// 0x17
	println(a.toString())
	// hi!
	a.test()
	a.value()
	a.pointer()
	fmt.Printf("v:%p,%v\n", &a, a)
}

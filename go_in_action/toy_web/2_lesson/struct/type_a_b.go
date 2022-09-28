package main

import "fmt"

func main() {
	fake := FakeFish{}
	// fake 无法调用原来 Fish 的方法
	//fake.Swim()
	fake.FakeSwim()

	// 转换成Fish
	td := Fish(fake)
	// 真的成了Fish
	td.Swim()

	cake := CakeFish{}
	// 这里调用自己的方法
	cake.Swim()

	td = Fish(cake)
	// 真的变成了鱼
	td.Swim()
}

type Fish struct {
}

func (f Fish) Swim() {
	fmt.Printf("我是鱼，假装自己是一直鸭子\n")
}

// FakeFish 定义一个新的类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎\n")
}

// CakeFish 定义一个新的类型
type CakeFish Fish

func (f CakeFish) Swim() {
	fmt.Printf("我是蛋糕鱼，吱吱\n")
}

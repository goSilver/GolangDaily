package main

import "fmt"

func main() {
	// 指针用*表示
	var p *ToyDuck = &ToyDuck{}
	// 解引用，得到结构体
	var duck ToyDuck = *p
	duck.Swim()

	var nilDuck *ToyDuck
	if nilDuck == nil {
		fmt.Println("nil")
	}
}

package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck
	duck1 := &ToyDuck{}
	duck1.Swim()

	duck2 := ToyDuck{}
	duck2.Swim()

	// duck3 是 *ToyDuck
	duck3 := new(ToyDuck)
	duck3.Swim()

	// go会初始化零值
	var duck4 ToyDuck
	duck4.Swim()

	// 空指针，会panic
	//var duck5 *ToyDuck
	//duck5.Swim()

	// 初始化按字段赋值
	duck6 := ToyDuck{
		Color: "Red",
		Price: 23,
	}
	duck6.Swim()

	duck7 := ToyDuck{"Blue", 12}
	duck7.Swim()

	duck8 := ToyDuck{}
	duck8.Color = "black"
	duck8.Swim()
}

type ToyDuck struct {
	Color string
	Price uint64
}

func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一只鸭，我是%s，%d一只\n", t.Color, t.Price)
}

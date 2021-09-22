package main

import "fmt"

func main()  {
	// 声明变量
	var i int
	// 给i赋值
	i = 10
	// 打印
	fmt.Println("i=", i)

	// Golang变量使用的三种方式
	// 1、指定变量类型，声明后若不赋值，使用默认值
	var j int
	fmt.Println("j=", j)
	// 2、根据值自行判断变量类型（类型推导）
	var k = 10.11
	fmt.Println("k=", k)
	// 3、省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	name := "tom"
	fmt.Println("name=", name)
}
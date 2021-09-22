package main

import "fmt"

// 定义全局变量
var n4 = 44
var n5 = 55
var name2 = "jack"
// 上面的声明方式，也可以改成一次性声明
var (
	n6 = 66
	n7 = 77
	name3 = "allen"
)

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

	// 一次性声明多个变量
	n1, name, n3 := 100, "tom~", 12.34
	fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	// 全局变量声明
	fmt.Println("n4=", n4, "n5=", n5, "name2=", name2)
	fmt.Println("n6=", n6, "n7=", n7, "name3=", name3)

}
package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	// 当直接输出byte值，就是输出对应字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	// 如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c, c2=%c\n", c1, c2)

	// var c3 byte = '北'  overflow
	var c3 int = '北'
	fmt.Println("c3=%c, c3对应码值=%d", c3, c3)
}

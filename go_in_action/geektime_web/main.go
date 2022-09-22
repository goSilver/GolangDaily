package main

import "unicode/utf8"

func main() {
	println("hello world")

	// go len()计算的是字节长度
	println(len("你好"))                      // 6
	println(utf8.RuneCountInString("你好"))   // 2
	println(utf8.RuneCountInString("你好ab")) // 4
}

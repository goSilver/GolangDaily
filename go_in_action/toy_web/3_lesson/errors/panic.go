package main

import "fmt"

func main() {

	defer func() {
		if data := recover(); data != nil {
			fmt.Printf("hello, panic:%v\n", data)
		}
		fmt.Println("恢复后从这里继续执行")
	}()

	panic("Boom")
	fmt.Println("这里将不会执行")
}

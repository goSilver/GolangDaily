package main

import "fmt"

func main() {
	arr := make([]int, 0, 10)
	arr = append(arr, 3)
	fmt.Printf("arr:%+v", arr)
}

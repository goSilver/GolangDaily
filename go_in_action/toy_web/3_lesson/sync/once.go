package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

var once sync.Once

func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}

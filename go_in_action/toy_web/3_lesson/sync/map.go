package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("cat", "tom")
	m.Store("mouse", "jerry")

	val, ok := m.Load("cat")
	if ok {
		fmt.Println(len(val.(string)))
	}
}

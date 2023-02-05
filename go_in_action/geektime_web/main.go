package main

import "fmt"

func main() {

	var arr []string
	arr = append(arr, "A")
	arr = append(arr, "B")
	arr = append(arr, "C")

	for _, s := range arr {
		switch s {
		case "A":
			fallthrough
		case "B":
			fmt.Println("B")
		case "C":
			fmt.Println("C")
		default:
			fmt.Println("unknown")
		}
	}
}

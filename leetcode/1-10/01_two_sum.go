package main

import "fmt"

func main() {
	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 4}
	target := 6
	//res := twoSumBruteForce(nums, target)
	res := twoSumMap(nums, target)
	fmt.Printf("res:%+v", res)
}

// twoSumBruteForce 蛮力法，双重for循环
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// twoSumMap 巧用map
func twoSumMap(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		numMap[num] = i
	}
	for i := 0; i < len(nums); i++ {
		index, ok := numMap[target-nums[i]]
		if ok && index != i {
			return []int{i, index}
		}
	}
	return []int{-1, -1}
}

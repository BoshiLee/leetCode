package main

import "fmt"

func main() {
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
}

func TwoNumberSum(array []int, target int) []int {
	seen := map[int]int{}
	for _, num := range array {
		search := target - num
		_, ok := seen[search]
		if (ok) {
			return []int{num, seen[search]}
		}
		seen[num] = num
	}
	return []int{}
}
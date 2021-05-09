package main

import "fmt"

func main() {
	fmt.Println(IsValidSubsequence([]int{1, 1, 6, 1}, []int{1, 1, 1, 6}))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	seen := 0
	if len(sequence) > len(array) {
		return false
	}
	for _, search := range sequence {
		seen = contains(array[seen:], seen, search)
		if seen < 0 {
			return false
		}
	}
	return true
}

func contains(slice []int, offset int, item int) int {
	for i, s := range slice {
		if s == item {
			return i + offset + 1
		}
	}
	return -1
}

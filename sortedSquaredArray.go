package main

import "fmt"

func main() {
	fmt.Println(SortedSquaredArray([]int{1, 9, 49, 81, 484, 900}))
}

func SortedSquaredArray(array []int) []int {
	left := 0
	right := len(array) - 1
	i := len(array) / 2
	if len(array) % 2 == 1 {
		array[len(array) / 2] *= array[len(array) / 2]
	}
	for i > 0 {
		fmt.Println(i)
		array[left] *= array[left]
		array[right] *= array[right]
		if array[left] > array[right] {
			temp := array[right]
			array[right] = array[left]
			array[left] = temp
		}
		left ++
		right --
		i --
	}
	return array
}

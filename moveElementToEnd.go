package main

import "fmt"

func main() {
	fmt.Println(MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
}

func MoveElementToEnd(array []int, toMove int) []int {
	p := 0
	for i, val := range array {
		if val != toMove {
			array[i] = toMove
			array[p] = val
			p++
		}
	}
	return array
}

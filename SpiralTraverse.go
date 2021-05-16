package main

import "fmt"

func main() {
	//    fmt.Println(SpiralTraverse([][]int{
	//    {1, 2, 3, 4},
	//    {12, 13, 14, 5},
	//    {11, 16, 15, 6},
	//    {10, 9, 8, 7},
	//}))
	fmt.Println(SpiralTraverse([][]int{
		{27, 12, 35, 26},
		{25, 21, 94, 11},
		{19, 96, 43, 56},
		{55, 36, 10, 18},
		{96, 83, 31, 94},
		{93, 11, 90, 16},
	}))
}

func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	if len(array[0]) <= 0 {
		return []int{}
	}
	i := 0
	bottom := len(array) - 1
	right := len(array[0]) - 1
	top := 0
	left := 0
	vertical := true   // true is down false is top
	horizontal := true // true is right false is left
	var result []int
	for i < len(array)*len(array[0]) {
		if horizontal {
			for v := left; v <= right; v++ {
				result = append(result, array[top][v])
				i++
			}
			top++
			horizontal = !horizontal

		} else {
			for v := right; v >= left; v-- {
				result = append(result, array[bottom][v])
				i++
			}
			bottom--
			horizontal = !horizontal
		}
		if vertical {
			for v := top; v <= bottom; v++ {
				result = append(result, array[v][right])
				i++
			}
			right--
			vertical = !vertical
		} else {
			for v := bottom; v >= top && v >= 0; v-- {
				result = append(result, array[v][left])
				i++
			}
			left++
			vertical = !vertical
		}
	}
	return result
}

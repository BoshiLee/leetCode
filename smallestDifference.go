package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17}))
}

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	p1, p2 := 0, 0
	distancePair := []int{array1[p1], array2[p2]}
	for p1 < len(array1) && p2 < len(array2) {
		if array1[p1] == array2[p2] {
			return []int{array1[p1], array2[p2]}
		}
		distancePair = compare(array1[p1], array2[p2], distancePair)
		if array1[p1] < array2[p2] {
			p1++
		} else {
			p2++
		}
	}
	return distancePair
}

//func SmallestDifference(array1, array2 []int) []int {
//    var distancePair []int
//    for _, a1 := range array1 {
//        for _, a2 := range array2 {
//            distancePair = compare(a1, a2, distancePair)
//        }
//    }
//    return distancePair
//}
//
func compare(a1 int, a2 int, distancePair []int) []int {
	distance := 0
	if abs(a1) > abs(a2) {
		distance = a1 - a2
	} else {
		distance = a2 - a1
	}
	if len(distancePair) == 0 {
		return []int{a1, a2}
	}
	current := abs(distancePair[0]) - abs(distancePair[1])
	if abs(distance) < abs(current) {
		return []int{a1, a2}
	}
	return distancePair
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

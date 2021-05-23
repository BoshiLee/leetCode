package main

import (
    "fmt"
    "sort"
)

func main() {
	fmt.Println(MergeOverlappingIntervals(
		[][]int{
			{9, 12},
            {45, 90},
			{12, 54},
            {43, 49},
			{91, 93},
		}))
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	order := [][]int{}
    headOrder := []int{}
    headMap := map[int]int{}
	for i, interval := range intervals {
        headOrder = append(headOrder, interval[0])
        headMap[interval[0]] = i
	}
	sort.Ints(headOrder)
	for _, head := range headOrder {
	    order = append(order, intervals[headMap[head]])
    }
	moveIndex := 0
	result := [][]int{}
	for _, current := range order {
		if len(result) == 0 {
			result = append(result, current)
			continue
		}
		prev := result[moveIndex]
		if current[0] <= prev[1] { // overlap
		    if current[1] <= prev[1] {
		        result[moveIndex][1] = prev[1]
            }
            if current[1] > prev[1] {
                result[moveIndex][1] = current[1]
            }
			result[moveIndex][0] = prev[0]
		} else {
            result = append(result, current)
            moveIndex++
        }
	}
	return result
}

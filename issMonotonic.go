package main

import "fmt"

func main() {
	fmt.Println(IsMonotonic([]int{1, 2, 0}))
}

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	postive := array[0] > 0
	var increasing bool
	if postive {
		increasing = array[0] < array[1]
	} else {
		increasing = array[0] > array[1]
	}
	for i := 2; i+1 < len(array); i++ {
		if postive {
			if increasing && array[i] > array[i+1] {
				return false
			}
			if !increasing && array[i] < array[i+1] {
				return false
			}
		} else {
			if increasing && array[i] < array[i+1] {
				return false
			}
			if !increasing && array[i] > array[i+1] {
				return false
			}
		}
	}
	return true
}

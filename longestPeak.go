package main

import "fmt"

func main() {
	fmt.Println(LongestPeak([]int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}))
}

func LongestPeak(array []int) int {
	tips := map[int]int{}
	for i, val := range array {
		if i-1 >= 0 && i+1 < len(array) {
			pre := array[i-1]
			next := array[i+1]
			if pre < val && val > next {
				tips[val] = i
			}
		}
	}
	longPeak := 0
	peaks := map[int]int{longPeak: 0}
	for k, v := range tips {
		peaks[k] = 1
		for i := v; i > 0; i-- {
			if array[i] > array[i-1] {
				peaks[k] += 1
			} else {
				break
			}
		}
		for i := v; i < len(array) && i + 1 < len(array); i++ {
			if array[i] > array[i+1] {
				peaks[k] += 1
			} else {
				break
			}
		}
		if peaks[k] > peaks[longPeak] {
			longPeak = k
		}
	}
	return peaks[longPeak]
}

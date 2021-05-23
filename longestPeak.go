package main

import "fmt"

func main() {
	fmt.Println(LongestPeak([]int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}))
}

func LongestPeak(array []int) int {
	longPeak := 0
	peaks := map[int]int{longPeak: 0}
	for v, val := range array {
		if v-1 >= 0 && v+1 < len(array) {
			pre := array[v-1]
			next := array[v+1]
			if pre < val && val > next {
				peaks[val] = 1
				for i := v; i > 0; i-- {
					if array[i] > array[i-1] {
						peaks[val] += 1
					} else {
						break
					}
				}
				for i := v; i < len(array) && i+1 < len(array); i++ {
					if array[i] > array[i+1] {
						peaks[val] += 1
					} else {
						break
					}
				}
				if peaks[val] > peaks[longPeak] {
					longPeak = val
				}
			}
		}
	}
	return peaks[longPeak]
}

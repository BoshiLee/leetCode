package main

import "fmt"

func main() {
	fmt.Println(LongestPeak([]int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}))
}

func LongestPeak(array []int) int {
	increase := true
	peaks := map[int]int{}
	previousPeakEnd := 0
	longestLength := 0
	for i, val := range array {
		if i + 1 < len(array) {
			increase = val < array[i + 1]
			currentLength := i - previousPeakEnd + 1
			if increase {
				if val > array[i + 1] {
					_, ok := peaks[longestLength]
					if ok && currentLength - longestLength > longestLength{
						peaks[currentLength - longestLength] = i
					} else {
						peaks[currentLength] = i
					}
					longestLength = currentLength
					previousPeakEnd = i
					continue
				}
			}
			if !increase {
				if val < array[i + 1] {
					_, ok := peaks[longestLength]
					if ok && currentLength - longestLength > longestLength{
						peaks[currentLength - longestLength] = i
					} else {
						peaks[currentLength] = i
					}
					longestLength = currentLength
					previousPeakEnd = i
					continue
				}
			}
		}
	}
	return longestLength
}

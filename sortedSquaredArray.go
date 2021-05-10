package main

import "fmt"

func main() {
	fmt.Println(SortedSquaredArray([]int{-50, -13, -2, -1, 0, 0, 1, 1, 2, 3, 19, 20}))
	//fmt.Println(insert([]int{1,2}, 1, 0))
	//fmt.Println(SortedSquaredArray([]int{-1, -1, 2, 3, 3, 3, 4}))
}

func SortedSquaredArray(array []int) []int {
	var newArray []int
	for i:=0 ; i<len(array); i += 1 {
		val := array[i] * array[i]
		if len(newArray) == 0  {
			newArray = append([]int{val}, newArray...)
			continue
		}
		if val <= newArray[0] {
			newArray = insert(newArray, 0, val)
			continue
		}
		if val >= newArray[len(newArray) - 1] {
			newArray = append(newArray, val)
			continue
		}
		start := 0
		end := len(newArray) - 1
		for end >= start {
			middle := int(uint(end+start) >> 1)
			if val == newArray[middle] {
				newArray = insert(newArray, middle, val)
				break
			}
			 if val > newArray[middle] {
				start = middle + 1
			} else {
				end = middle - 1
			}

			if len(newArray[start:end]) == 1 {
				fmt.Println(newArray, val, middle)
				if  val < newArray[start] {
					newArray = insert(newArray, start, val)
					break
				}
				if  val > newArray[start] {
					newArray = insert(newArray, start + 1, val) // end
					break
				}
				if  val < newArray[end] {
					newArray = insert(newArray, end, val)
					break
				}
				if val > newArray[end] {
					newArray = insert(newArray, end + 1, val)
					break
				}

			}
		}
	}
	return newArray
}

func insert(a []int, index int, value int) []int {
	a = append(a[:index+1], a[index:]...) // Step 1+2
	a[index] = value                      // Step 3
	return a
}
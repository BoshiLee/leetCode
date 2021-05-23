package main

func FirstDuplicateValue(array []int) int {
	record := map[int]int{}
	for _, val := range array {
		if value, exist := record[val]; exist {
			return value
		} else {
			record[val] = val
		}
	}
	return -1
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(NonConstructibleChange([]int{5, 7, 1, 1, 2, 3, 22}))
	//fmt.Println(NonConstructibleChange([]int{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4}))
}

func NonConstructibleChange(coins []int) int {
	if len(coins) == 1 && coins[0] > 1 || len(coins) == 0 {
		return 1
	}
	chage := 0
	sort.Ints(coins)
	for i, val := range coins {
		chage += val
		if i+1 < len(coins) && chage+1 < coins[i+1] {
			return chage + 1
		}
	}
	return chage + 1
}

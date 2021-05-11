package main

import (
    "fmt"
    "sort"
)

func main() {
    fmt.Println(ThreeNumberSum([]int{
        12, 3, 1, 2, -6, 5, -8, 6,
    }, 0))
}

//func ThreeNumberSum(array []int, target int) [][]int {
//    seen := map[int]int{}
//    var result = make([][]int, 0)
//    sort.Ints(array)
//    for i, num := range array {
//        for j := i+1 ; j < len(array) ; j++ {
//            search := target - num - array[j]
//            _, ok := seen[search]
//            if ok {
//                r := []int{num, search, array[j] }
//                result = append(result, r)
//            }
//            seen[array[j]] = array[j]
//        }
//        seen = map[int]int{}
//    }
//    return result
//}
func ThreeNumberSum(array []int, target int) [][]int {
    var result = make([][]int, 0)
    sort.Ints(array)
    for i := 0; i< len(array) - 2; i ++ {
        left := i+1
        right := len(array) -1
        for left < right {
            sum := array[left] + array[i] + array[right]
            if sum == target {
                result = append(result, []int{array[i], array[left], array[right]})
                left += 1
                right -= 1
            }
            if sum < target {
                left += 1
            }
            if sum > target {
                right -= 1
            }
        }
    }
    return result
}
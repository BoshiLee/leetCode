package main

func main() {

}

func ArrayOfProducts(array []int) []int {
    var result []int
    for i, _ := range array {
        current := 1
        for j, val := range array {
            if i == j {
                continue
            }
            current *= val
        }
        result = append(result, current)
    }
    return result
}
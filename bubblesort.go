package main
import (
    "fmt"
)
var arr []int = []int{1,3,2,4,8,6,7,2,3,0,9,8,10,11,12}
func bubbleSort(input []int) {
    n := len(input)
    swapped := true
    for swapped {
        swapped = false
        for i := 1; i < n; i++ {
            if input[i-1] > input[i] {
                input[i], input[i-1] = input[i-1], input[i]
                swapped = true
            }
        }
    }
    fmt.Println(input)
}
func main() {
    fmt.Println("Hello")
    bubbleSort(arr)
}

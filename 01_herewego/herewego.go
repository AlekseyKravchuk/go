// https://www.youtube.com/watch?v=CF9S4QZuV30 - learn the Go programming language in one video
// 14:12 / 52:57
// Каждая программа в go начинается с объявления package
package main

import "fmt"

func main() {
	size := 5
	// var nums [size]float64{10.64, 34.11, 29.98948, 2.9, 4.5}
	nums := [...]float64{10.64, 34.11, 29.98948, 2.9, 4.5}

	for i := 0; i < size; i++ {
		fmt.Printf("nums[%d] = %.1f \n", i, nums[i])
	}
}

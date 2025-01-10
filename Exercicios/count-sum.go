package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	var result [2]int
	var count int
	var negatives int
	for _, number := range numbers {
		if number > 0 {
			count += 1
		} else if number < 0 {
			negatives += number
		}
	}
	result[0] = count
	result[1] = negatives
	fmt.Println(result)
}

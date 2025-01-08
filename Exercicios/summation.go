package main

import "fmt"

func main() {
	var sum int
	number := 8
	for i := 1; i <= number; i++ {
		sum += i
	}
	fmt.Println(sum)
}

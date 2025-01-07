package main

import "fmt"

func main() {
	num := 1
	if num < 0 {
		fmt.Println(num)
	} else if num > 0 {
		fmt.Println(num * (-1))
	} else {
		fmt.Println(num)
	}
}

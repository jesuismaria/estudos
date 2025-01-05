package main

import "fmt"

func main() {
	fmt.Println(soma(42, 13))

	sub := subtracao(10, 5)
	fmt.Println(sub)
}

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

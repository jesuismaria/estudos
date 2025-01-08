package main

import "fmt"

func main() {
	var sum int
	numbers := []int{1, -4, 7, 12}
	for _, num := range numbers {
		if num > 0 {
			sum = sum + num
		}
	}
	fmt.Println(sum)
}

// Exercício "Sum of positive" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

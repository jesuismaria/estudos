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

// Exercício "Grasshopper-Summation" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

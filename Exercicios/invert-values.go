package main

import "fmt"

func main() {
	num := -5
	if num < 0 {
		fmt.Println(num)
	} else {
		fmt.Println(num * -1)
	}
}

// Exercício "Invert Values" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

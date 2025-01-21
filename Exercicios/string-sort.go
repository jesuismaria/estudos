package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "CodeWars"
	var evenLetters string
	var oddLetters string

	letters := strings.Split(str, "")
	for i, letter := range letters {
		if i%2 == 0 {
			evenLetters += letter
		} else {
			oddLetters += letter
		}

	}
	fmt.Println(evenLetters, " ", oddLetters)
}

// Exercício "Odd-Even String Sort" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

//Minha resolução

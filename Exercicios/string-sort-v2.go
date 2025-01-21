package main

import "fmt"

func main() {
	word := "CodeWars"
	var evenLetters, oddLetters string

	for i := 0; i < len(word); i++ {
		if i%2 == 0 {
			evenLetters += string(word[i])
			continue
		}
		oddLetters += string(word[i])
	}
	fmt.Println(fmt.Sprintf("%s %s", evenLetters, oddLetters))
}

// Exercício "Odd-Even String Sort" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

//Resolução da prof

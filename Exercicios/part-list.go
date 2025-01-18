package main

import (
	"fmt"
	"strings"
)

func main() {
	elements := []string{"az", "toto", "picaro", "zone", "kiwi"}
	var finalElements string

	for i := 1; i < len(elements); i++ {
		element := fmt.Sprintf("(%s, %s)", strings.Join(elements[:i], " "), strings.Join(elements[i:], " "))
		finalElements += element
	}
	fmt.Println(finalElements)

}

// Exercício "Parts of a list" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

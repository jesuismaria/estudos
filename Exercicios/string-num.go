package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "605"
	num, _ := strconv.Atoi(str)
	fmt.Println(num)
}

// Exercício "Convert string to a number" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

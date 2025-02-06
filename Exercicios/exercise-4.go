package main

import "fmt"

type teste int

var x teste

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

}

// Exercício do curso de Go
// da prof Ellen disponibilizado no youtube

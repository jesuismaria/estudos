package main

import "fmt"

type teste int

var x teste
var y int //int - tipo subjacente de Teste

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x) //conversao
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}

// Exercício do curso de Go
// da prof Ellen disponibilizado no youtube

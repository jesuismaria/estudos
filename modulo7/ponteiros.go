package main

import "fmt"

// Ponteiros: uma variável que ao invés de armazenar
// um valor (1, "olá", true ...) armazena uma memória

func zeroValue(i int) {
	i = 0
}

func zeroPointer(i *int) {
	*i = 0
}
func main() {
	//Memória -> essa memória tem um endereço -> endereço guarad valor

	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória:", &i)
	//&variavel -> aponta para o endereço de memória da variável

	zeroValue(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	a := &i
	fmt.Println("variável a:", a)
	fmt.Println("Variável *a:", *a)
	// usa o * em relação ao endereço! Não do valor
	fmt.Println("variável &a:", &a)
}

// Aula 7.3 Ponteiros
// by Steph Cardoso

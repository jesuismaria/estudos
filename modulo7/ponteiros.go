package main

import "fmt"

// Ponteiros: uma variável que ao invés de armazenar
// um valor (1, "olá", true ...) armazena uma memória

func main() {
	//Memória -> essa memória tem um endereço -> endereço guarad valor

	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória:", &i)
	//&variavel -> aponta para o endereço de memória da variável
}

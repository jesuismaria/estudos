package main

import "fmt"

func main() {
	// Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	// fmt.Println(array[0])
	// fmt.Println(array[1])
	// fmt.Println(array[0], array[1])
	// fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	// pega o 2, 3 e 5
	fmt.Println(numPrimos[0:3])
	// pega tudo antes da posição 1
	fmt.Println(numPrimos[:1])
	// pega tudo depois da posição 1
	fmt.Println(numPrimos[1:])




// Array

// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do array
// Por conta do tamanho fixo, não é tão usado. Só em casos específicos

// Slice

// Tipo o array, mas sem tamanho fixo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len() retorna o tamanho do slice
// Função append() usada para adicionar valores

package main

import "fmt"

// IF / ELSE
// SE / SE NÃO

func main() {

	valor := 1
	//if (condição) { <ação> }
	if valor == 1 { //true
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor difere de 1, valor é :", valor)
	}

	if valor == 1 {
		fmt.Println("Valor é igual a 1")
	} else if numero == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor é diferente de 1 e 2")
	}

	// par ou ímpar
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}
	//OU
	numero := 7
	if numero%2 == 0 {
		fmt.Printf("%d é par", numero)
	} else {
		fmt.Printf("%d é ímpar", numero)
	}

}

// OPERAÇÕES
// Soma: 1 + 1
// Subtração: 2 - 1
// Divisão: 10 / 2
// Multiplicação: 2 * 2
// Resto da divisão por x: 7%2 (resto da divisão por 2)

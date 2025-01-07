package main

import (
	"fmt"
	"time"
)

// LOOPS
// Laços de repetição
// Repetir tarefas

func main() {

	// i++ -> soma 1
	// i-- -> subtrai 1
	sum := 0
	for 1 := 0; i < 10; i++ {
		sum += i
		// é a mesma coisa que: sum = sum + i
		// sum -= i -> sum = sum - i
		// É como se no final do loop adiconasse 1 ao índice
	}
	fmt.Println(sum)

	// LOOP INFINITO
	for {
		fmt.Println("Golang do zero")
		time.Sleep(2 * time.Second)
	}

	// FOR RANGE
	frutas := []string{"laranja", "maça", "banana", "kiwi"}
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}
}

// OBS: Quando cria variável dentro do escopo For
//ela só existe lá dentro

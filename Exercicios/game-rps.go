package main

import "fmt"

func main() {
	p1 := "paper"
	p2 := "scissors"

	if p1 == "scissors" && p2 == "paper" {
		fmt.Println("Player 1 won!")
	}
	if p1 == "rock" && p2 == "scissors" {
		fmt.Println("Player 1 won!")
	}
	if p1 == "paper" && p2 == "rock" {
		fmt.Println("Player 1 won!")
	}
	if p1 == "scissors" && p2 == "rock" {
		fmt.Println("Player 2 won!")
	}
	if p1 == "rock" && p2 == "paper" {
		fmt.Println("Player 2 won!")
	}
	if p1 == "paper" && p2 == "scissors" {
		fmt.Println("Player 2 won!")
	}
	if p2 == "scissors" && p1 == "rock" {
		fmt.Println("Player 2 won!")
	}
	if p2 == "rock" && p1 == "paper" {
		fmt.Println("Player 2 won!")
	}
	if p2 == "paper" && p1 == "scissors" {
		fmt.Println("Player 2 won!")
	}

	if p1 == p2 {
		fmt.Println("Draw!")
	}

}

// Exercício "Rock Paper Scissors" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

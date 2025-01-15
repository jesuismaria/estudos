package main

import (
	"fmt"
	"strings"
)

func main() {
	var points int
	games := []string{"3:1", "2:2", "0:1"}

	for _, game := range games {
		score := strings.Split(game, ":")

		if score[0] == score[1] {
			points += 1
		}
		if score[0] > score[1] {
			points += 3
		}
	}
	fmt.Println(points)
}

// Exercício "Total amount of points" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

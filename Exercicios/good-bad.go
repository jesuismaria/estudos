package main

import "fmt"

func main() {
	var sumIdeas int
	ideas := [5]string{"bad", "bad", "bad"}
	for _, idea := range ideas {
		if idea == "good" {
			sumIdeas += 1
		}
	}
	if sumIdeas == 1 || sumIdeas == 2 {
		fmt.Println("Publish!")
	} else if sumIdeas > 2 {
		fmt.Println("I smell a series!")
	} else {
		fmt.Println("Fail!")
	}

}

// Exercício "Well of ideas - Easy Version" do site Codewars
// indicado no curso de Go da Comuniade Dev Completo
// by Steph Cardoso

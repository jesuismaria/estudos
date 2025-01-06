package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["maria"] = 21
	idade["bento"] = 4
	fmt.Println(idade)
	fmt.Println(idade["maria"])
	fmt.Println(idade["bento"])

	anoNasc := map[string]int{
		"maria": 2003,
		"bento": 2008,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["maria"])
	fmt.Println(anoNasc["bento"])
	// add novo elemento ao map
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)
}

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro tipo
//map[k]v -> k = chave, v = valor

//map[string]int
//{ "maria": 21, "bento": 4}
// map[string]string
// { "maria":"barreto", "bento": "cardoso"}

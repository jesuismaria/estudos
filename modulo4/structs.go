package main

import "fmt"

// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes

// type <nome da estrutura> struct {<campos>}
type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	fmt.Println(Pessoa{"maria", 21})
	// Mais aconselhável setar os campos
	fmt.Println(Pessoa{Nome: "Bento", Idade: 2})
	fmt.Println(Pessoa{Nome: "Maria"})

	p1 := Pessoa{Nome: "Bob", Idade: 2}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 3
	fmt.Println(p1.Idade)

	p2 := Pessoa(Nome: "Patrik", Idade: 2)

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// structs + map
	alunos := map[string][]Pessoa{}
	alunos["programação"] = pessoas
	fmt.Println(alunos)

	var alunos = map[string][]Pessoa{
		"Programação": {{Nome: "Maria", Idade: 21}}
	}
}
//OBS: se voce criar uma struct e não setar todos os campos, retornará o zero value
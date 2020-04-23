package main

import "fmt"

type Pessoa struct{
	nome string
	telefone string
	idade  int
}

func main(){
	//instant
	var p Pessoa

	p.nome="Joao"
	p.telefone="12345678"
	p.idade=30
	fmt.Printf("nome= %s\n", p.nome)
	fmt.Printf("telefone=%s\n", p.telefone)
	fmt.Printf("idade=%d\n", p.idade)
}

package main

import (
	"fmt"
)

func main() {
	nome, idade := nomeIdade()
	fmt.Println("Meu nome é", nome, "e tenho", idade, "anos")
}

func nomeIdade() (string, int) {
	nome := "Ricardo"
	idade := 51
	return nome, idade
}
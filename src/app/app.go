package main

import (
	"fmt"
	"os"
)

func main() {
	
	introducao()

	opcoes()

	comando := commando()

	switch comando {
	case 1:
		monitoramento()
	case 2: 
		fmt.Println("Exibindo Logs...")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

	fmt.Println("")
}

func introducao() {
	nome := "Fahham"
	versao := 1.0
	
	fmt.Println("----------------------------------------------")
	fmt.Println(" Olá", nome)
	fmt.Println(" A versão desta app é", versao)
	fmt.Println("----------------------------------------------")
}

func opcoes() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
	fmt.Println("----------------------------------------------")
}

func commando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)
	return comando
}

func monitoramento() {
	fmt.Println("Monitorando...")
}

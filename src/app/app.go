package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
)

const vezes = 5
const delay = 5

func main() {

	introducao()

	for {

		opcoes()
		
		comando := commando()

		switch comando {
		case 1:
			monitoramento()
			fmt.Println("")
		case 2: 
			fmt.Println("Exibindo Logs...")
		case 3:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
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
	fmt.Println("")
	fmt.Println("Digite uma opção...")
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
	fmt.Println("______________")
	fmt.Println("")

	sites := []string{
		"https://report.apps.tsuru.gcp.i.globo",
		"https://www.globo.com",
	}

	for i := 0; i < vezes ; i++ {
		fmt.Println("Tentativa:", i)
		for i, site := range sites {
			fmt.Println("testando o site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	// fmt.Println(resp)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!!!")
	} else {
		fmt.Println("Site", site, "não está respondendo, Status Code:", resp.StatusCode )
	} 
}
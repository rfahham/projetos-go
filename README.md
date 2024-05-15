# GO

Linguagem de programação criada pelo google em 2006.

## Instalando
    https://golang.org/dl/
    Fazer o Download do respectivo Sistema Operacional

## Saber a versão
    go version
    go version go1.22.2 linux/amd64

## Criando o Go Workspace
    mkdir go
    cd go
    mkdir pkg - Onde ficam os pacotes compartilhados das aplicações (o Go é uma linguagem bastante modular)
    mkdir src - Onde ficam o sorce code
    mkdir bin - Onde ficam os binários (compilados)

## Criando o primeiro script
Criar a pasta do projeto
    scr/
    mkdir hello
    cd hello

## Criar o arquivo hello.go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }

## Compilando o código
    go build <nome do script>

## Executando o código
    ./hello

## Compilando e Executando o script
    go run hello.go


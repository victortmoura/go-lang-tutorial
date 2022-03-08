package main

import (
	"fmt"
	"reflect"
)

func main() {

	// definindo tipo da variavel
	// var nome string = "Victor Moura"
	// var nome = "Victor Moura"

	// inferencia do tipo da variavel atraves do :=
	nome := "Victor Moura"
	versao := 1.1

	fmt.Println("Prazer", nome)
	fmt.Println("Você está na versão:", versao)
	fmt.Println("* Tipo da variável versão é:", reflect.TypeOf(versao))
	fmt.Println()

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	// ao utilizar a funcao ScanF, deve-se inferir o modificador %d (que representa digito)
	// fmt.Scanf("%d", comando)

	// Lembrar de colocar o '&' pois o input ira ser salvo no endereco de memoria da variavel
	fmt.Scan(&comando)

	fmt.Println("O comando selecionado foi", comando)
}

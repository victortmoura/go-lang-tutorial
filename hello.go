package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// for sem nada ira rodar indefinidamente
	for {
		exibirIntroducao()
		exibirMenu()

		comando := lerComando()

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
			monitorar()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando ", comando, "inexistente")
			os.Exit(-1)
		}
	}

	// Outro modo de realizar a validacao do input do usuario
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa")
	// } else {
	// 	fmt.Println("Comando ", comando, "inexistente")
	// }
}

func exibirIntroducao() {
	// definindo tipo da variavel
	// var nome string = "Victor Moura"
	// var nome = "Victor Moura"

	// inferencia do tipo da variavel atraves do :=
	nome := "Victor Moura"
	versao := 1.1

	fmt.Println("Prazer", nome)
	fmt.Println("Você está na versão:", versao)
	fmt.Println()
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	// ao utilizar a funcao ScanF, deve-se inferir o modificador %d (que representa digito)
	// fmt.Scanf("%d", comando)

	// Lembrar de colocar o '&' pois o input ira ser salvo no endereco de memoria da variavel
	fmt.Scan(&comandoLido)

	return comandoLido
}

func monitorar() {
	url := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("Site", url, "foi carregado com sucesso !")
	} else {
		fmt.Println("Site", url, "está com problemas. Status Code:", resp.StatusCode)
	}

}

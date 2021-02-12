package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()

	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida...")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	fmt.Println("Seja bem vindo!")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("Menu:")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Progrma")
}

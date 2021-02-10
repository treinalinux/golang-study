package main

import "fmt"

func main() {
	nome := "Alan"
	version := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Esse programa está na versão", version, "\n")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Progrma \n")

	var comando int
	fmt.Scan(&comando)
	//fmt.Scanf("%d", &comando)

	fmt.Println("\nO endereço na memória é:", &comando)
	fmt.Println("Comando:", comando)
}

package main

import "fmt"

func main() {
	fmt.Println("Menu:")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Progrma")

	var comando int
	fmt.Scan(&comando)

	// if comando == 1 {
	// if 	fmt.Println("Monitorando...")
	// if } else if comando == 2 {
	// if 	fmt.Println("Exibindo Logs...")
	// if } else if comando == 0 {
	// if 	fmt.Println("Saindo do programa...")
	// if } else {
	// if 	fmt.Println("Opção inválida!")
	// if }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Opção inválida...")
	}
}

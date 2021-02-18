package main

import "fmt"

func main() {
	fmt.Println("Uso do defer:")
	fmt.Println("Abrindo a conexão com banco de dados...")
	// O primeiro defer sempre será executado por ultimo
	defer fmt.Println(" 1 defer - Estou fechando a conexão com banco de dados")
	defer fmt.Println(" 2 defer - Estou fechando a conexão com banco de dados")
	executaConsulta()
}

// executaConsulta
func executaConsulta() {
	fmt.Println("Conectando no banco de daos")
}

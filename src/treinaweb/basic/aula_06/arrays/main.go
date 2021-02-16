package main

import (
	"fmt"
)

func main() {
	fmt.Println("Criando Arrays")
	var amigos [5]string
	fmt.Println(amigos)
	// Uso do len()
	for i := 0; i < len(amigos); i++ {
		// %s é porque recebe uma string
		fmt.Scanf("%s", &amigos[i])
	}
	fmt.Println(amigos)
	fmt.Println("Seus amigos são: ")

	// For no estilo foreach
	for _, amigo := range amigos {
		fmt.Printf(" - %s \n", amigo)
	}
	fmt.Println("")

	arrayInicializado := [3]int{2, 3, 6}
	fmt.Println(arrayInicializado)
	fmt.Println("")

	var matriz [3][3]int
	// for aninhado
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Digite um número: ")
			fmt.Scanf("%d", &matriz[i][j])
		}
	}
	fmt.Println(matriz)
}

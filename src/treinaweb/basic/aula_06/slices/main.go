package main

import (
	"fmt"
)

func main() {
	fmt.Println("Use slices no lugar de arrays")
	// Criando um slices com o make([]tipo, quantidade)
	// A quantidade definida não vai limitar o tamanho quando usar append()
	amigos := make([]string, 3)
	nome := ""
	// for no estilo while
	i := 0
	for nome != "q" {
		// Recebendo a entrada do usuário
		fmt.Print("Digite o nome do amigo ou 'q' para sair: ")
		fmt.Scanf("%s", &nome)
		if nome != "q" {
			if i < 3 {
				amigos[i] = nome
			} else {
				amigos = append(amigos, nome)
			}
			i++
		}
	}
	fmt.Println("")
	fmt.Println(amigos)
	fmt.Println("")
	fmt.Printf("Você tem %d amigos \n", len(amigos))
	for _, nomeAmigo := range amigos {
		fmt.Printf(" - %s \n", nomeAmigo)
	}
	// outro exemplo, pegando somente a posição 1 e 2 (vai usar o 3, pois ele não entra)
	fmt.Println("")
	outroSlice := amigos[1:3]
	fmt.Println(outroSlice)
	// Mais um
	fmt.Println("")
	maisUmSlice := amigos[:3]
	fmt.Println(maisUmSlice)
}

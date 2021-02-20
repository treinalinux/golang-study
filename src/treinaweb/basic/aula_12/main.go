package main

import (
	"fmt"

	estrutura "treinalinux.com/alanalves/carros"
)

func main() {
	fmt.Println("Trabalhando com interfaces")
	fmt.Println(".................")

	carro := estrutura.Carro{Modelo: "Palio", Marca: "Fiat"}

	opcao := 0
	for opcao != 3 {
		fmt.Println("O que você desaja fazer? ")
		fmt.Println(" 1 - Acelerar ")
		fmt.Println(" 2 - Frear ")
		fmt.Println(" 3 - Sair")
		fmt.Print(".........: ")
		fmt.Scanf("%d", &opcao)
		var err error = nil
		switch opcao {
		case 1:
			/*
				err = carro.acelerar()
				Devemos passar a referencia, nesse caso não passar o *ponteiro porque não está na memória
				somento o carro está na memória.
			*/
			err = fazerAcelerar(&carro)
		case 2:
			// err = carro.frear()
			err = fazerFrear(&carro)
		}
		if err != nil {
			fmt.Printf(" ** Não foi possível executar a ação: %s ** \n", err)
		} else if opcao != 3 {
			fmt.Printf("O carro %s da marca %s está a %.2f km/h \n", carro.Modelo, carro.Marca, carro.Velocidade)
		}
	}
	fmt.Println("Fim da execução!")
}

// fazerAcelerar criado para interface Acelerador
func fazerAcelerar(veiculo estrutura.Acelerador) error {
	return veiculo.Acelerar()
}

// fazerFrear criado para interface Acelerador
func fazerFrear(veiculo estrutura.Freador) error {
	return veiculo.Frear()
}

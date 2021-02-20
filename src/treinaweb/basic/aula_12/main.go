package main

import (
	"fmt"

	estrutura "treinalinux.com/alanalves/carros"
)

/*
	O código atual foi refatorado para a criação do pacote em:
	/home/alan/go/src/treinalinux.com/alanalves/carros/estruturas.go
	Aqui no código main ele foi chamado com:

	apelido   caminho
	estrutura "treinalinux.com/alanalves/carros"

	Outra obsevação é que antes o arquivo estruturas.go em estrutura ...
	teve o build e o install:
	/home/alan/go/src/treinalinux.com/alanalves/carros $ go build
	/home/alan/go/src/treinalinux.com/alanalves/carros $ go install

	Com isso um link carros.a foi criado no caminho abaixo e pode ser chamado aqui.
	/home/alan/go/pkg/linux_amd64/treinalinux.com/alanalves/carros.a

*/
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

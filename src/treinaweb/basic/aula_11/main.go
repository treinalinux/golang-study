package main

import (
	"errors"
	"fmt"
)

// Acelerador nova interface criada
type Acelerador interface {
	acelerar() error
}

// Freador nova interface criada
type Freador interface {
	frear() error
}

// Carro é uma struct em outras linguagens seria uma classe
type Carro struct {
	modelo     string
	marca      string
	velocidade float32
}

// carro chamado na definiçao do ponteiro (carro *Carro) no metodo acelerar
func (carro *Carro) acelerar() error {
	if carro.velocidade < 15 {
		carro.velocidade += 5
		return nil

	} else {
		// error strings should not be capitalized or end with punctuation or a newline
		return errors.New("o carro já está em sua velocidade máxima")
	}
}

// carro chamado na definiçao do ponteiro (carro *Carro) no metodo frear
func (carro *Carro) frear() error {
	if carro.velocidade > 0 {
		carro.velocidade -= 5
		return nil

	} else {
		return errors.New("o carro já está parado")
	}
}

func main() {
	fmt.Println("Trabalhando com interfaces")
	fmt.Println(".................")

	carro := Carro{modelo: "Palio", marca: "Fiat"}

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
			fmt.Printf("O carro %s da marca %s está a %.2f km/h \n", carro.modelo, carro.marca, carro.velocidade)
		}
	}
	fmt.Println("Fim da execução!")
}

// fazerAcelerar criado para interface Acelerador
func fazerAcelerar(veiculo Acelerador) error {
	return veiculo.acelerar()
}

// fazerFrear criado para interface Acelerador
func fazerFrear(veiculo Freador) error {
	return veiculo.frear()
}

package carros

import "errors"

// Acelerador nova interface criada
type Acelerador interface {
	Acelerar() error
}

// Freador nova interface criada
type Freador interface {
	Frear() error
}

// Carro é uma struct em outras linguagens seria uma classe
type Carro struct {
	Modelo     string
	Marca      string
	Velocidade float32
}

// Acelerar chamado na definiçao do ponteiro (carro *Carro) no metodo acelerar
func (carro *Carro) Acelerar() error {
	if carro.Velocidade < 15 {
		carro.Velocidade += 5
		return nil

	}
	return errors.New("o carro já está em sua velocidade máxima")
}

// Frear chamado na definiçao do ponteiro (carro *Carro) no metodo frear
func (carro *Carro) Frear() error {
	if carro.Velocidade > 0 {
		carro.Velocidade -= 5
		return nil

	}
	return errors.New("o carro já está parado")
}

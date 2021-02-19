package main

import (
	"fmt"
)

type Buzinador interface {
	buzinar() string
}

type Carro struct {
	marca string
}

func (carro *Carro) buzinar() string {
	return " - Buzina: FonFon"
}

func main() {
	carro := Carro{}
	fmt.Println(chamarBuzina(&carro))
}

func chamarBuzina(buzinador Buzinador) string {
	return buzinador.buzinar()
}

package main

import "fmt"

// ContaCorrente foi criada para struct da conta corrente
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// Sacar c de conta que estiver chamado o saque
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

func main() {
	contaDoAlan := ContaCorrente{}
	contaDoAlan.titular = "Alan"
	contaDoAlan.saldo = 500

	fmt.Println(contaDoAlan.saldo)

	fmt.Println(contaDoAlan.Sacar(501))

	fmt.Println(contaDoAlan.saldo)
}

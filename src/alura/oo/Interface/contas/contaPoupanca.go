package contas

import "alura/oo/Interface/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

// Sacar c de conta que estiver chamado o saque
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente!"
}

// Depositar criado para depositar na conta
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso! R$", c.saldo
	}
	return "Deposito falhou, pois é menor do que zero! R$", c.saldo
}

// ObterSaldo saldo da conta poupança
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

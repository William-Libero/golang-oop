package contas

import "github.com/William-Libero/golang-oop/cliente"

type ContaPoupanca struct {
	Titular                              cliente.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) string {
	podeSacar := c.saldo >= valor && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		return "Depósito realizado com sucesso! Saldo da conta:", valor
	} else {
		return "Valor de depósito insuficiente Saldo da conta:", valor
	}
}

func (contaRemetente *ContaPoupanca) Transferencia(valor float64, contaDestino *ContaPoupanca) string {
	podeTransferir := contaRemetente.saldo >= valor && valor > 0
	if podeTransferir {
		contaRemetente.saldo -= valor
		contaDestino.Depositar(valor)
		return "Transferencia realizado com sucesso!"
	} else {
		return "Valor de transferencia insuficiente."
	}
}

func (conta *ContaPoupanca) GetSaldo() float64 {
	return conta.saldo
}

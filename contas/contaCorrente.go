package contas

import "github.com/William-Libero/golang-oop/cliente"

type ContaCorrente struct {
	Titular       cliente.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := c.saldo >= valor && valor > 0
	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		return "Depósito realizado com sucesso! Saldo da conta:", valor
	} else {
		return "Valor de depósito insuficiente Saldo da conta:", valor
	}
}

func (contaRemetente *ContaCorrente) Transferencia(valor float64, contaDestino *ContaCorrente) string {
	podeTransferir := contaRemetente.saldo >= valor && valor > 0
	if podeTransferir {
		contaRemetente.saldo -= valor
		contaDestino.Depositar(valor)
		return "Transferencia realizado com sucesso!"
	} else {
		return "Valor de transferencia insuficiente."
	}
}

func (conta *ContaCorrente) GetSaldo() float64 {
	return conta.saldo
}

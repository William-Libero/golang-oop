package main

import (
	"fmt"

	"github.com/William-Libero/golang-oop/cliente"

	"github.com/William-Libero/golang-oop/contas"
)

func main() {
	clienteWilliam := cliente.Titular{Nome: "William", CPF: "449.887.328-94", Profissao: "Desenvolvedor"}
	contaPoupancaDoWilliam := contas.ContaPoupanca{Titular: clienteWilliam, NumeroAgencia: 589, NumeroConta: 123456}

	contaDoWilliam := contas.ContaCorrente{Titular: clienteWilliam, NumeroAgencia: 589, NumeroConta: 123456}
	contaDoWilliam.Depositar(500)
	PagarBoleto(&contaDoWilliam, 50)
	fmt.Println(contaDoWilliam)
	fmt.Println(contaDoWilliam.GetSaldo())

	fmt.Println(contaPoupancaDoWilliam)
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

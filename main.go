package main

import (
	"fmt"

	"go/go-oop/cliente"

	"./contas"
)

func main() {
	contaDoWilliam := contas.ContaCorrente{Titular: cliente.Titular{"William", "449.887.328-94", "Desenvolvedor"}, NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}

	fmt.Println(contaDoWilliam)
}

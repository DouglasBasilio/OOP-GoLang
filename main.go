package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	c1 := contas.ContaCorrente{}
	c1.Depositar(100)
	PagarBoleto(&c1, 60)

	fmt.Println(c1.ObterSaldo())

	c2 := contas.ContaPoupanca{}
	c2.Depositar(500)
	PagarBoleto(&c2, 100)

	fmt.Println(c2.ObterSaldo())

}

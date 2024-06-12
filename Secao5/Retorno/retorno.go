package main

import (
	"fmt"
)

func main() {
	soma, subtracao, divisao, multiplicacao := operacoes(1, 2)
	fmt.Println(soma, subtracao, divisao, multiplicacao)
}

func operacoes(num1, num2 int) (
	soma int, subtracao int, divisao int, multiplicacao int) {

	soma = num1 + num2
	subtracao = num1 - num2
	divisao = num1 / num2
	multiplicacao = num1 * num2
	return
}

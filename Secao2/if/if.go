package main

import "fmt"

func main() {
	bonus := 100.0
	salario := 950.0
	var salarioMaisBonus float64 = salario
	if salario < 1000 {
		salarioMaisBonus += bonus
	}
	fmt.Println(salarioMaisBonus)

	var ehCarro bool = false
	var valorAuto = 1000.00
	if ehCarro {
		valorAuto += 55.50
	} else {
		valorAuto += 20.5
	}
	fmt.Println("Valor final: ", valorAuto)
}

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
}

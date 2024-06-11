package main

import "fmt"

func main() {
	salario := 900.00
	tipo := "CLT"

	if salario < 1000 && tipo == "CLT" {
		salario -= (salario * 0.08)
	} else if salario < 1000 && tipo == "PJ" {
		salario -= (salario * 0.05)
	} else if salario <= 1200 {
		salario -= (salario * 0.10)
	} else {
		salario -= (salario * 0.15)
	}
	fmt.Println("Salario Final: ", salario)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Teste de lista encadeada")
	var lista []int
	lista = append(lista, 4, 9, 6, 7)
	fmt.Println("Lista: ", lista)
	fmt.Println("Lista pos 1: ", lista[0])
	fmt.Println("Lista pos 2: ", lista[1])
	fmt.Println("Lista pos 3: ", lista[2])
	fmt.Println("Tamanho da lista: ", len(lista))

	lista2 := make([]int, 1)
	lista2[0] = 10
	lista2 = append(lista2, 1)
	lista2 = append(lista2, 4)
	fmt.Println(reflect.TypeOf(lista2))

	var somaTotal int
	for i := 0; i < len(lista2); i++ {
		somaTotal += lista2[i]
	}

	fmt.Println(somaTotal / len(lista2))
}

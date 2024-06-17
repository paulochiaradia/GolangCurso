package main

import (
	"fmt"
	"math"
)

func main() {
	var lista []interface{}
	lista = append(lista, 2)
	lista = append(lista, "Teste")
	lista = append(lista, []int{1, 5, 6, 7})
	lista = append(lista, math.Pi)

	for _, valor := range lista {
		if v, ok := valor.(string); ok {
			fmt.Println(v + " string")
		} else {
			fmt.Println(valor)
		}

	}

}

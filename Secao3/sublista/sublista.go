package main

import "fmt"

func main() {
	listaCompleta := []int{2, 10, 9, 4, 5, 8, 1, 3}
	var subLista = make([]int, 0)
	for i := 0; i < len(listaCompleta); i++ {
		if listaCompleta[i] < 5 {
			subLista = append(subLista, listaCompleta[i])
			fmt.Println("Posicao [", i, "]", subLista)
		}

	}
}

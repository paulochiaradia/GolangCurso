package main

import "fmt"

func main() {
	primeiraLista := []int{1, 2, 3, 5, 6, 7, 8, 9}
	segundaLista := primeiraLista[:3]
	terceiraLista := primeiraLista[4:]
	ultimoItem := primeiraLista[len(primeiraLista)-1:]

	fmt.Println(primeiraLista)
	fmt.Println(segundaLista)
	fmt.Println(terceiraLista)
	fmt.Println(ultimoItem)

}

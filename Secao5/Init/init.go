package main

import "fmt"

func main() {
	ola("Nivaldo")
}

func ola(nome string) {
	fmt.Println("Ola " + nome)
}

func init() {
	fmt.Println("Iniciando o Programa")
}

func init() {
	fmt.Println("Iniciando o programa 2")
}

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	texto := "palavra"
	tamanho := len(texto)
	j := 0
	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
	}
	fmt.Println()
	for j < tamanho {
		fmt.Println(string(texto[j]))
		j++
	}
}

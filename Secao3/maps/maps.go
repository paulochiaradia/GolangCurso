package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["sp"] = 100
	m["rj"] = 90
	m["cg"] = 50
	fmt.Println(m["rj"])

	if valor, existe := m["mg"]; existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Nao existe")
		fmt.Println(existe)
	}
	delete(m, "cg")
	for chave, valor := range m {
		fmt.Println("Cidade:", chave, "Habitantes: ", valor)
	}
}

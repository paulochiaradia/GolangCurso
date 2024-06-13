package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["sp"] = 100
	m["rj"] = 90
	fmt.Println(m["rj"])

	valor, existe := m["mg"]
	if existe {
		fmt.Println(valor)
	} else {
		fmt.Println("Nao existe")
		fmt.Println(existe)
	}
}

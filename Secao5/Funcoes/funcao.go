package main

import "fmt"

func main() {
	imprimirMensagem("Asdrubal")
}

func imprimirMensagem(mensagem string) {
	mensagem += ", bom dia"
	fmt.Println(mensagem)
}

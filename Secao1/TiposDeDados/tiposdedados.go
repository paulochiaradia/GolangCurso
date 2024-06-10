package main

import "fmt"

func main() {
	//Variavel nao pode alterar seu tipo durante o programa
	texto := "Texto1"
	//Incorreto texto=10
	texto = "Texto2"
	fmt.Println(texto)
	numero := 10
	medida := 5.5754
	ehMasculino := false

	fmt.Println(numero)
	fmt.Println(medida)
	fmt.Println(ehMasculino)

}

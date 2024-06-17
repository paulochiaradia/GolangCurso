package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("c:/texto.txt")
	if err != nil {
		panic(err)
	}
	//unreached code

	fmt.Println("Teste")
}

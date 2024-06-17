package main

import (
	"fmt"
	"os"
)

func ShowText() {
	fmt.Println("Manipulacao finalizada")
}

func main() {
	file, err := os.Create("./settings.txt")

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	defer ShowText()
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte("Teste"))
	if err != nil {
		panic(err)
	}
}

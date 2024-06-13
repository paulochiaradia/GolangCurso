package model

import "fmt"

type Automovel struct {
	Ano    int
	Modelo string
	Cor    string
}

func (a Automovel) Ligar() {
	fmt.Println("Automovel esta ligado")
}

type Carro struct {
	Automovel
	Cavalos int
}

type Moto struct {
	Automovel
	Cilindrada int
}

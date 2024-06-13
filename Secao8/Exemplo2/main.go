package main

import (
	"fmt"
	"paulochiaradia/GolangCurso/Secao8/Exemplo2/model"
)

func main() {

	carro := model.Carro{
		Automovel: model.Automovel{Ano: 2024, Modelo: "Nivus", Cor: "Branco"},
		Cavalos:   125,
	}
	carro.Ligar()
	fmt.Println(carro.Modelo)
}

package main

import (
	"fmt"
	model2 "paulochiaradia/GolangCurso/Secao8/Exemplo1/model"
	"time"
)

func main() {
	endereco := model2.Endereco{
		Rua:    "Rua x",
		Numero: 154,
		Cidade: "sp",
	}
	fmt.Println(endereco)

	pessoa1 := model2.Pessoa{
		Nome:             "Euclesimar",
		Endereco:         endereco,
		DataDeNascimento: time.Date(2000, 12, 17, 0, 0, 0, 0, time.Local),
	}

	pessoa1.CalculaIdade()
	fmt.Println(pessoa1.Idade)
}

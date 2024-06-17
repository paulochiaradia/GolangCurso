package main

import (
	"fmt"
	"paulochiaradia/GolangCurso/Secao8/Exercicio/model"
	"time"
)

func main() {
	var itens []model.ItemDaCompra
	itens = append(itens, model.ItemDaCompra{Item: "Arroz"})
	itens = append(itens, model.ItemDaCompra{Item: "Feijao"})
	itens = append(itens, model.ItemDaCompra{Item: "Carne"})
	fmt.Println(itens)

	compra := model.Compra{
		Data:    time.Now(),
		Mercado: "Mercado Astolfo",
		Itens:   itens,
	}

	fmt.Println(compra)
	var novaCompra model.Compra
	novaCompra.InicializaCompra()
	fmt.Println(novaCompra)

}

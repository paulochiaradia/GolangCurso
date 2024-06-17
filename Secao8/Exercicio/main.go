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
	var compraVazia model.Compra
	compraVazia.InicializaCompra()
	fmt.Println(compraVazia)

	novaCompra, err := model.NovaCompra(time.Now(), "AAA", []string{"banana", "pera", "uva"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(*novaCompra)
	}

}

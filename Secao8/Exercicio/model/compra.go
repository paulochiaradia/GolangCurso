package model

import (
	"errors"
	"time"
)

type Compra struct {
	Data    time.Time
	Mercado string
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Item string
}

func (c *Compra) InicializaCompra() *Compra {
	return &Compra{
		Data:    time.Time{},
		Mercado: "",
		Itens:   make([]ItemDaCompra, 0),
	}
}

func NovaCompra(data time.Time, mercado string, itensDaCompra []string) (*Compra, error) {
	if mercado == "" {
		return nil, errors.New("mercado e obrigatorio")
	}

	if len(itensDaCompra) == 0 {
		return nil, errors.New("os itens sao obrigatorios")
	}

	var itens []ItemDaCompra

	for _, valor := range itensDaCompra {
		itens = append(itens, ItemDaCompra{Item: valor})
	}
	return &Compra{
		Data:    data,
		Mercado: mercado,
		Itens:   itens,
	}, nil
}

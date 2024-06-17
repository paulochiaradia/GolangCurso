package model

import (
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

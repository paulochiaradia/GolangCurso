package model

import "time"

type Compra struct {
	Data    time.Time
	Mercado string
	Itens   []ItemDaCompra
}

type ItemDaCompra struct {
	Item string
}

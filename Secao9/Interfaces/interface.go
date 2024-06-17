package main

import (
	"errors"
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}
type retangulo struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func ImprimeArea(g geometria) {
	fmt.Println(g.area())
}

func main() {
	circulo := circulo{
		raio: 5,
	}
	retangulo := retangulo{
		altura:  10,
		largura: 5,
	}
	ImprimeArea(circulo)
	ImprimeArea(retangulo)
	ExibeErro(errors.New("aaaa"))
	p := problemaDeNetwork{
		rede:     true,
		hardware: false,
	}
	ExibeErro(p)
}

type problemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p problemaDeNetwork) Error() string {
	if p.rede {
		return "Problema de rede"
	} else if p.hardware {
		return "Problema de hardware"
	} else {
		return "Outro Problema"
	}
}

func ExibeErro(err error) {
	fmt.Println(err.Error())
}

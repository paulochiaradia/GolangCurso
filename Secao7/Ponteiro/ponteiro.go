package main

import "fmt"

func main() {
	x := 5
	y := &x
	fmt.Println(x, y)
	*y = 50
	fmt.Println(x, *y)
	alteraValor(&x)
	fmt.Println(&x, *y)
}

func alteraValor(valor *int) {
	*valor = 10
}

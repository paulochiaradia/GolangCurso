package main

import (
	"fmt"
	"reflect"
)

func main() {
	num1 := 10.0
	num2 := 20.0
	//Operadores + - * /
	result := num1 + num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

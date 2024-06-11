package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var numeroInt8 int8 = 127
	var numeroInt int = int(numeroInt8)
	fmt.Println(numeroInt)
	b, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Erro")
	}
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
}

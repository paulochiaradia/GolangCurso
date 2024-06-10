package main

import (
	"fmt"
	"reflect"
)

func main() {
	const taxa = 4
	fmt.Println(taxa)
	fmt.Println(reflect.TypeOf(taxa))
}

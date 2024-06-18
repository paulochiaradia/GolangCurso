package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 10000; i++ {
		go showMessage(strconv.Itoa(i))
	}
}

func showMessage(message string) {
	fmt.Println(message)
}

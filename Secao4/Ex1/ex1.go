package main

import "fmt"

func main() {
	somaTotal := 0
	arr := []int{10, 5}
	//arr := [2]int{}

	for _, valor := range arr {
		somaTotal += valor
	}
	fmt.Println(somaTotal)
}

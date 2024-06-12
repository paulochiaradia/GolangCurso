package main

import "fmt"

func main() {
	arr := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	var somaMenor, somaMaior int
	for _, valor := range arr {
		if valor <= 5 {
			somaMenor += valor
		} else {
			somaMaior += valor
		}
	}
	fmt.Println(somaMaior)
	fmt.Println(somaMenor)
}

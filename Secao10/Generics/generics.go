package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"a", "jj", "ff"}
	newSlice1 := reverse(slice1)
	newSlice2 := reverse(slice2)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)
}

func reverse[T int | string](slice []T) []T {
	newSlice := make([]T, len(slice))
	newSliceLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newSlice[newSliceLen] = slice[i]
		newSliceLen--
	}
	return newSlice
}

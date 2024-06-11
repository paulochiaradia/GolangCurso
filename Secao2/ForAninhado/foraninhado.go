package main

import "fmt"

func main() {
	for numBase := 1; numBase <= 10; numBase++ {
		for numMult := 1; numMult <= 10; numMult++ {
			fmt.Println(numBase, " x ", numMult, " = ", numBase*numMult)
		}
	}
}

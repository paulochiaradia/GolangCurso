package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go callDatabase(&wg)
	go callInternal(&wg)
	go callApi(&wg)
	wg.Wait()
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado Database")
	wg.Done()
}

func callInternal(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado internal")
	wg.Done()
}

func callApi(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finalizado Api")
	wg.Done()
}

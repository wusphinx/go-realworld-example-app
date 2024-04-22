package main

import (
	"fmt"
	"sync"
)

var (
	toOdd  = make(chan struct{})
	toEven = make(chan struct{})
)

func main() {
	sg := &sync.WaitGroup{}
	a := 10
	sg.Add(2)
	go printOdd(a, sg)
	go printEven(a, sg)
	sg.Wait()
}

func printOdd(high int, sg *sync.WaitGroup) {
	defer sg.Done()

	for i := 1; i <= high; i += 2 {
		if i != 1 {
			<-toOdd
		}
		fmt.Println(i)
		toEven <- struct{}{}
	}
}

func printEven(high int, sg *sync.WaitGroup) {
	defer sg.Done()

	for i := 2; i <= high; i += 2 {
		<-toEven
		fmt.Println(i)
		if i != high {
			toOdd <- struct{}{}
		}
	}
}

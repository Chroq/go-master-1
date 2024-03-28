package main

import (
	"fmt"
	"time"
)

func takeCoffees(ch chan<- int, coffeNumber int) {
	ch <- coffeNumber
}

func deliverCoffees(ch <-chan int) {
	for i := range ch {
		fmt.Println("Coffe number: ", i)
	}
}

func main() {
	ch := make(chan int, 2)

	go takeCoffees(ch, 1)
	go takeCoffees(ch, 2)
	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()

	deliverCoffees(ch)
}

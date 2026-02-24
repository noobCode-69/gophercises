package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pizza struct {
	orderNumber int
}

func makePizza(n int) Pizza {
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	return Pizza{orderNumber: n}
}

func worker(out chan Pizza, quit chan bool) {
	for i := 1; ; i++ {
		pizza := makePizza(i)
		select {
		case out <- pizza:
		case <-quit:
			fmt.Println("Worker stopping...")
			close(out)
			return
		}
	}
}

func main() {
	pizzaCh := make(chan Pizza)
	quitCh := make(chan bool)
	go worker(pizzaCh, quitCh)
	for p := range pizzaCh {
		fmt.Println("Got pizza", p.orderNumber)
		if p.orderNumber == 5 {
			quitCh <- true
		}
	}
}
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup;

func printSomething(s string) {
	defer wg.Done()
	fmt.Println(s)
}	

func main(){
	wg.Add(1)
	go printSomething("This is the first thing to be printed")
	wg.Wait()
	fmt.Println("This is the second thing to be printed")
}
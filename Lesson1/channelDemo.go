package main

import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("producing data: %d\n", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for k := range ch {
		fmt.Printf("Get data: %d\n", k)
	}
}

func main() {
	ch := make(chan int, 2)
	go produce(ch)
	consumer(ch)
}

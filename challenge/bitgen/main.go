package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch0 := make(chan int)
	go create0(ch0)
	go create1(ch1)
	for i := 1; i <= 100000; i++ {
		select {
		case v := <-ch1:
			fmt.Print(v)
		case v := <-ch0:
			fmt.Print(v)
		}
	}
}

func create1(ch chan<- int) {
	for {
		ch <- 1
	}
}

func create0(ch chan<- int) {
	for {
		ch <- 0
	}
}
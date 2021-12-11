package main

import "fmt"

func sum(x, y int, c chan int) {
	s := x + y
	c <- s
}

func main() {
	s := make(chan int)
	a, b := 1,2
	go sum(a, b, s)
	fmt.Printf("%v + %v = %v", a, b, <-s)
}
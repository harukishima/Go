package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go sendData(&c)
	go receiveData(&c)
	time.Sleep(1e9)
}

func sendData(c *chan int) {
	for i := 0; i <= 10; i++ {
		*c <- i
	}
}

func receiveData(c *chan int) {
	for {
		fmt.Println(<-*c)
	}
}
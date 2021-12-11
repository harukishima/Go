package main

import (
	"fmt"
)

func badCall() {
  a, b := 10, 0
  n := a / b  
  fmt.Println(n)
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()
	badCall()
}

func main() {
	test()
}
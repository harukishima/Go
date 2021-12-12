package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := setInterval(func() {
		fmt.Println("Tick")
	}, 1e9)
	time.Sleep(5e9)
	ticker.Stop()
	_ = setTimeout(func() {
		fmt.Println("Timeout")
	}, 2e9)
	time.Sleep(3e9)
}

func setInterval(fc func(), tick int) *time.Ticker {
	ticker := time.NewTicker(time.Duration(tick))
	go func() {
		for _ = range ticker.C {
			fc()
		}
	}()
	return ticker
}

func setTimeout(fc func(), timeout int) *time.Timer {
	timer := time.NewTimer(time.Duration(timeout))
	go func() {
		<-timer.C
		fc()
	}()
	return timer
}

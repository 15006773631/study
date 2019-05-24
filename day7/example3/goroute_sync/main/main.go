package main

import (
	"fmt"
)

func send(ch chan int, exitchan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitchan <- a
}
func recv(ch chan int, exitchan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitchan <- a
}
func main() {
	ch := make(chan int, 10)
	exitchan := make(chan struct{}, 2)
	go send(ch, exitchan)
	go recv(ch, exitchan)
	total := 0
	for _ = range exitchan {
		total++
		if total == 2 {
			break
		}
	}
}

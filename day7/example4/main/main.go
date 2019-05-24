package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	go func() {
		var i int
		for {
			ch <- i
			time.Sleep(time.Second)
			i++
		}
	}()
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}

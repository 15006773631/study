package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 100000)
	for i := 0; i < 100000; i++ {
		ch <- i
	}
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}

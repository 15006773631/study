package main

import "fmt"

func add() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
func main() {
	f := add()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}

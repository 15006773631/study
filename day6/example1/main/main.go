package main

import "fmt"

type Test interface {
}

func main() {
	var a interface{}
	var b int
	a = b
	fmt.Printf("type of a %T\n", a)
}

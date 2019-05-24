package main

import "fmt"

type Example interface {
	GetName() string
	Run()
}

func main() {
	var a interface{}
	var b int
	var c float32
	a = b
	a = c
	fmt.Println()
}

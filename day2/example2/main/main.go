package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Intn(10)
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println(a)
	}
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println(a)
	}
}

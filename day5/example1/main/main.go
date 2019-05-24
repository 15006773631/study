package main

import (
	"fmt"
)

type student struct {
	Age   int
	Name  string
	score int
}

func main() {
	var stu student
	stu.Age = 18
	stu.Name = "tyler"
	stu.score = 99
	fmt.Println("u name is ", stu.Name)
	var stu1 = student{
		Age:  20,
		Name: "tyler",
	}
	fmt.Println(stu1.Age)
	fmt.Println(stu1.Name)
}

package main

type Student struct {
	Name  string
	Age   int
	score float32
	left  *Student
	right *Student
}

func main() {
	var root *Student = new(Student)
	root.Name = "tyler"
	root.Age = 18
	root.score = 99
}

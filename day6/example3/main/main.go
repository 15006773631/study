package main

import "fmt"

type Read interface {
	Read()
}
type Write interface {
	Write()
}

type ReadWrite interface {
	Read
	Write
}

type file struct {
}

func (f *file) Read() {
	fmt.Println("read data")
}
func (f *file) Write() {
	fmt.Println("Write data")
}
func Test(rw ReadWrite) {
	rw.Read()
	rw.Write()
}
func main() {
	var f file
	Test(&f)
}

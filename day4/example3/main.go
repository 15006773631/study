package main

import "fmt"

func ssort(a []int) {
	for i := 0; i < len(a); i++ {
		var min = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

}
func main() {
	b := [...]int{5, 2, 6, 7, 8, 4, 123, 1234, 22}
	ssort(b[:])
	fmt.Println(b)
}

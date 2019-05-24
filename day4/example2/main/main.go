package main

import "fmt"

func bsort(a []int64) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

}
func main() {
	b := [...]int64{5, 2, 6, 7, 8, 4, 123, 1234, 22}
	bsort(b[:])
	fmt.Println(b)
}

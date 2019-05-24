package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	rand.Seed(time.Now().Unix())
	n = rand.Intn(101)
	fmt.Println("请输入一个1到100的数字：")
	for {
		var a int
		fmt.Scanf("%d\n", &a)
		var b = false
		switch {
		case a == n:
			fmt.Println("恭喜你猜对了")
			b = true

		case a > n:
			fmt.Println("猜大了")
		case a < n:
			fmt.Println("猜小了")
		}
		if b {
			break
		}
	}

}

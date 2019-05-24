package main

import (
	"fmt"
	"net/http"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {
	for _, v := range url {
		resp, err := http.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}
		fmt.Printf("head succ,status:%v\n", resp.Status)
	}
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var Mytemplate *template.Template

type person struct {
	Name  string
	Age   int
	Title string
}

func userinfo(w http.ResponseWriter, r *http.Request) {
	var arr []person
	p := person{Name: "Mary001", Age: 10, Title: "我的个人网站"}
	p1 := person{Name: "Mary002", Age: 12, Title: "我的个人网站"}
	p2 := person{Name: "Mary003", Age: 13, Title: "我的个人网站"}
	arr = append(arr, p)
	arr = append(arr, p1)
	arr = append(arr, p2)
	Mytemplate.Execute(w, arr)
}
func Intetemplate(filename string) (err error) {
	Mytemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}
func main() {
	Intetemplate("C:/Users/1040317/go/src/study/day8/http_template/index.html")

	http.HandleFunc("/usr/info", userinfo)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

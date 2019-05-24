package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string `json:"username""`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Email    string `json:email`
}

func jsontest() {
	user1 := &user{
		Username: "tyler",
		Age:      18,
		Sex:      "man",
		Email:    "1370990924@qq.com",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal is failed err:%s", err)
	}
	fmt.Printf("json:%s", string(data))
}
func main() {
	jsontest()
}

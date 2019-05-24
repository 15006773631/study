package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username string
	Age      int
	Sex      string
	Email    string
}

func main() {
	user1 := &user{
		Username: "tyler",
		Age:      18,
		Sex:      "man",
		Email:    "1370990924@qq.com",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.marshal failed, err:", err)
		return
	}
	fmt.Printf(string(data))
}

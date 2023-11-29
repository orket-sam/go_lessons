package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	IsVerified bool   `json:"verified"`
	Email      string `json:"email"`
}

func main() {

	jsonString := `{"name":"Orket","age":24,"verified":false,"email":"samirorket@gmail.com"}`
	var user user
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		println("oops")
	}
	fmt.Printf("%T", user)
}

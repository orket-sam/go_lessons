package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name        string `json:"name"`
	Yob         int
	Nationality string
}

func main() {
	user := User{}
	userJson, _ := json.Marshal(user)
	fmt.Println(string(userJson))

	userString := `{"name":"orket","Yob":0,"Nationality":"Kenyan"}`

	var user2 User
	err := json.Unmarshal([]byte(userString), &user2)
	if err == nil {
		fmt.Println(user2)
	}
}

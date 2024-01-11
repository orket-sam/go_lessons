package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	var user User

	stringUser := `{"name":"orket","age":23,"i":5616}`

	if err := json.Unmarshal([]byte(stringUser), &user); err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name)

}

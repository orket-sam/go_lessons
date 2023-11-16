package main

import "fmt"

func main() {
	users := []struct {
		name string
		age  int
	}{{"shanaya", 5}}

	fmt.Println(users)
}

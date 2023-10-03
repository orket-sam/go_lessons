package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHello(10)
}

func sayHello(n int) {
	for i := 1; i < n; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

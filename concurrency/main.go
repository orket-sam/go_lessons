package main

import (
	"fmt"
	"time"
)

func heavyTask(n int, c chan string) {

	c <- fmt.Sprintf("task %d has began \n", n)

	time.Sleep(time.Duration(n) * time.Second)

	c <- fmt.Sprintf("%d has completed \n", n)
}

func main() {
	c := make(chan string)

	for i := 0; i < 5; i++ {
		go heavyTask(i, c)
	}

	result := make([]string, 10)

	for i := range result {
		result[i] = <-c
		fmt.Println(result[i])

	}

}

package greetings

import "fmt"

func greetings(name string) string {

	message := fmt.Sprintln("Hi, Welcome!", name)
	return message
}

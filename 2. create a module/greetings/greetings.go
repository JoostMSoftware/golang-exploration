package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hey, %v. Welcome to this tutorial!", name)
	return message
}

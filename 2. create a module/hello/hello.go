package main

import (
	"fmt"

	"joostmsoftware.com/greetings"
)

func main() {
	message := greetings.Hello("Joost")
	fmt.Println(message)

}

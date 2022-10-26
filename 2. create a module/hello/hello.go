package main

import (
	"fmt"
	"log"

	"joostmsoftware.com/greetings"
)

func main() {
	/*
		Set properties of the predefined Loggers, including
		the log entry prefix and add a flag to disable printing
		the time, source file and line number.
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("Test")

	/*
		If an error was returned, print it to the console and
		exit the program.
	*/
	if err != nil {
		log.Fatal(err)
	}

	/*
		If no error was returned, print the returned message
		to the console
	*/
	fmt.Println(message)
}

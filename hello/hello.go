package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	//A slice of names
	names := []string{" Ruchi ", " Priya ", " Gaurav "}

	//Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	//Request greeting
	//message, err := greetings.Hello("Ruchi")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	//if there was no error, print the
	//returned message to console

	// Get a greeting message and print it
	fmt.Println()
	fmt.Println(messages)
	fmt.Println()
}

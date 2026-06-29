package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Request a greeting message.
	//message, err := greetings.Hello("Golang")

	names := []string{"Rohit", "Mohit", "Rajat"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(message)
	fmt.Println(messages)
}

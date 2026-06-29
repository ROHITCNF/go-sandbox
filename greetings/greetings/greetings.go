package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello returns a greeting for the named person.

func Hello(name string) (string, error) {
	//return error if name is empty
	if name == "" {
		return "", errors.New("empty name")
	}
	//Return a greeting that embeds the name in a message.
	var message string
	message = fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// randomFormat can only be accessible in own package
func randomFormat() string {
	//slice of message formats
	formats := []string{
		"Hi %v . Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// return a random index of slices
	return formats[rand.Intn(len(formats))]
}

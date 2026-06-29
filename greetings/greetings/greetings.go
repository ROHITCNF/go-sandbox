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

// hellos function returns a map with names and
// messages for each name

func Hellos(names []string) (map[string]string, error) {
	//initialize a map  which can hold key bas string and value has string
	// make === new in JS
	messages := make(map[string]string)

	//loop through the received names call Hello function for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// no error means we have got our greeting messages
		messages[name] = message
	}
	return messages, nil

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

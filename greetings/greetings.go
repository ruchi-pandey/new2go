package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	//if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	//if a name was received, return value embeding the name
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	//Map for associating names and messages
	messages := make(map[string]string)
	//Loop through the slice of names, calling
	//message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.

func randomFormat() string {
	//A slice of message formats
	formats := []string{
		"Hey there, %v!",
		"Hail, %v! Namaste",
		"Konichiwa %v, how have you been?",
	}

	//Return randomly selected message format
	//splice the strings is indexed per format
	return formats[rand.Intn(len(formats))]

}

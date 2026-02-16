package greetings

import (
	"fmt"
	"errors"
	// "log"
	"math/rand"
)

// Helle returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Param name is null or empty")
	}
	// message := fmt.Sprintf("Hi, %v. Welcome to the world of Go!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		// log.Print(_)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
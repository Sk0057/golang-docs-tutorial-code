package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//If not name, then error
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomGreetingFormat(), name)
	return message, nil
}

// randomGreetingFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomGreetingFormat() string {
	// A slice of message greetingFormats.
	greetingFormats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return greetingFormats[rand.Intn(len(greetingFormats))]
}

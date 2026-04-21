package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// Capitalized name when is a exported package function
func Hello(input string) (string, error) {

	if strings.TrimSpace(input) == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), input)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	greets := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		greets[name] = message
	}

	return greets, nil
}

// lowecase name unexported package function
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Greate to see you, %v",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

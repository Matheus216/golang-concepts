package greetings

import "fmt"

func Hello(input string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", input)
	return message
}

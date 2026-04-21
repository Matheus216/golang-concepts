package greetings

import (
	"strings"
	"testing"
)

func TestHelloMessageSuccess(t *testing.T) {
	nameToTest := "Fernandinho"
	response := Hello(nameToTest)

	if !strings.Contains(response, "Welcome") {
		t.Errorf("Don't have Welcome, in this result: %v", response)
	}
}

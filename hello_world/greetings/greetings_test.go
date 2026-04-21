package greetings

import (
	"fmt"
	"strings"
	"testing"
)

func TestHelloMessageSuccess(t *testing.T) {
	nameToTest := "Fernandinho"
	response, error := Hello(nameToTest)

	if error != nil {
		t.Errorf("Invalid test case, this error happens: %v", error.Error())
	}

	if !strings.Contains(response, "Welcome") &&
		!strings.Contains(response, "Greate") &&
		!strings.Contains(response, "Hail") {
		t.Errorf("Don't have Welcome, in this result: %v", response)
	}
}

func TestWhenPassInvalidNameShouldReturnError(t *testing.T) {
	response, error := Hello("")

	if response != "" {
		t.Error("Text should be empty")
	}

	if error == nil {
		t.Error("Error should be filled")
	}

	fmt.Printf("error: %v\n", error.Error())
}

func TestWhenPassSpacesAsParametersShouldReturnError(t *testing.T) {
	response, error := Hello(" ")

	if error == nil || response != "" {
		t.Error("Should return error.")
	}
}

func TestWhenCallTwiceShouldReturnDifferenceMessage(t *testing.T) {
	responseF, err := Hello("Fernandinho")

	if err != nil {
		t.Error("Invalid fail point")
	}

	responseT, err := Hello("Fernandinho")

	if responseF == responseT {
		t.Error("Should be differences messages between the calls")
	}
}

func TestWhenSendMultiNamesShouldReturnMapString(t *testing.T) {
	request := []string{
		"Fernandinho",
		"Thiaguinho",
		"Rogeria",
		"Sebastiana",
	}

	response, err := Hellos(request)

	if err != nil || len(response) == 0 || response == nil {
		t.Error("Invalid return")
	}

	for _, name := range request {
		if response[name] == "" {
			t.Errorf("%v: name not found in the response", name)
		}
		fmt.Printf("name: %v, message:%v\n", name, response[name])
	}
}

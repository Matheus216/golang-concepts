package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Printf("Starting the project: %v \n\n", os.Args[0])

	names := []string{}

	for index, parameter := range os.Args {
		if index == 0 {
			continue
		}

		names = append(names, parameter)
	}

	greet, err := greetings.Hellos(names)

	if err == nil {
		fmt.Println(greet)
		return
	}

	log.Fatalln(err)
}

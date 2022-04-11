package main

import (
	"errors"
	"fmt"
)

func main() {
	var userInput1 string
	var userInput2 string

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Did you miss out entering your first or last name?")
		}
	}()

	fmt.Printf("\nPlease enter your first name: ")
	fmt.Scanln(&userInput1)
	fName := userInput1

	fmt.Printf("Please enter your last namme: ")
	fmt.Scanln(&userInput2)
	lName := userInput2

	if fName == "" {
		panic(errors.New("Runtime error: First name cannot be nil"))
	} else if lName == "" {
		panic(errors.New("Runtime error: Last name cannot be nil"))
	}

	fmt.Printf("Hello, %s %s!! Nice to meet you\n", fName, lName)
}

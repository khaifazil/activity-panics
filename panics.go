package main

import (
	"errors"
	"fmt"
	"strconv"
)

func printAllGames(){
	for i, game := range boardGames {
		fmt.Printf("%v : %s\n", i+1, game)
	}
}

func checkUserInput(result int) error {

	if result > len(boardGames) || result < 0 {
		return errors.New("Invalid selection!")
	}
	return nil
}

var boardGames = []string{"Carcassone", "Wildlife Safari", "Civilization"}
var userInput string

func main() {
	var msg string
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T) \n", err, err)
		}
	}()

	fmt.Println("The boardGames are : \n")

	printAllGames()

	fmt.Println("What is your favorite game?")
	fmt.Scanln(&userInput)

	result, _ := (strconv.Atoi(userInput) - 1)



	fmt.Println("Oh I see. So your favorite game is:", boardGames[result])
}
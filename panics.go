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

func checkUserInput(result int) {
	if result >= len(boardGames) || result < 0 {
		panic(errors.New("Invalid selection!"))
	}
}

var boardGames = []string{"Carcassone", "Wildlife Safari", "Civilization"}

func main() {
	var userInput string

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T) \n", err, err)
		}
	}()

	fmt.Println("The boardGames are : \n")

	printAllGames()

	fmt.Println("What is your favorite game?")
	fmt.Scanln(&userInput)

	userInputInt, _ := strconv.Atoi(userInput)
	result := userInputInt - 1
	checkUserInput(result)


	fmt.Println("Oh I see. So your favorite game is:", boardGames[result])
}